//go:build linux

package platforms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/bomkz/patchman/global"
	"golang.org/x/sys/unix"
)

var (
	WriteMsg = make(chan global.Message)
	ReadMsg  = make(chan global.Message)
	pongMsg  = make(chan bool)
)

const socketPath = "/tmp/patchman.sock"

func initParent() {
	parentPID := os.Getpid()

	// Clean up any leftover socket from a previous run
	os.Remove(socketPath)

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		global.FatalError(fmt.Errorf("failed to listen on unix socket: %w", err))
	}
	// Only owner can connect
	if err := os.Chmod(socketPath, 0600); err != nil {
		global.FatalError(fmt.Errorf("failed to chmod socket: %w", err))
	}

	childPID := launchHelper()
	allowedPIDs := map[int]bool{
		parentPID: true,
		childPID:  true,
	}

	conn := acceptAndVerify(l, allowedPIDs)

	go handleRead(conn)
	go handleWrite(conn)
	go heartBeat()
}

func initHelper() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var conn net.Conn
	for {
		var err error
		d := net.Dialer{}
		conn, err = d.DialContext(ctx, "unix", socketPath)
		if err == nil {
			break
		}
		if ctx.Err() != nil {
			writeLogPriviledged(err, "Error at: ConnectToParent:S1")
			return
		}
		time.Sleep(50 * time.Millisecond)
	}

	go handleRead(conn)
	go handleWrite(conn)
	go heartBeat()
}

func launchHelper() int {
	exe, err := os.Executable()
	if err != nil {
		global.FatalError(fmt.Errorf("could not resolve executable path: %w", err))
	}

	// Try pkexec first (graphical sudo prompt), fall back to sudo
	var cmd *exec.Cmd
	if _, err := exec.LookPath("pkexec"); err == nil {
		cmd = exec.Command("pkexec", exe)
	} else {
		cmd = exec.Command("sudo", "-n", exe)
	}

	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		global.FatalError(fmt.Errorf("failed to launch helper: %w", err))
	}

	return cmd.Process.Pid
}

func acceptAndVerify(l net.Listener, allowedPIDs map[int]bool) net.Conn {
	conn, err := l.Accept()
	if err != nil {
		global.FatalError(fmt.Errorf("accept failed: %w", err))
		return nil
	}

	unixConn, ok := conn.(*net.UnixConn)
	if !ok {
		conn.Close()
		global.FatalError(errors.New("connection is not a unix socket"))
		return nil
	}

	rawConn, err := unixConn.SyscallConn()
	if err != nil {
		conn.Close()
		global.FatalError(fmt.Errorf("could not get raw conn: %w", err))
		return nil
	}

	var cred *unix.Ucred
	var credErr error
	rawConn.Control(func(fd uintptr) {
		cred, credErr = unix.GetsockoptUcred(int(fd), unix.SOL_SOCKET, unix.SO_PEERCRED)
	})
	if credErr != nil {
		conn.Close()
		global.FatalError(fmt.Errorf("SO_PEERCRED failed: %w", credErr))
		return nil
	}

	if !allowedPIDs[int(cred.Pid)] {
		conn.Close()
		global.FatalError(fmt.Errorf("unauthorized PID: %d", cred.Pid))
		return nil
	}

	return conn
}

func handleRead(conn net.Conn) {
	decoder := json.NewDecoder(conn)

	for {
		var response global.Message
		if err := decoder.Decode(&response); err != nil {
			writeLogPriviledged(err, "Error at handleRead:S1")
			os.Exit(1)
		}
		switch response.Type {
		case "ping":
			WriteMsg <- global.Message{Type: "pong"}
		case "pong":
			pongMsg <- true
		default:
			ReadMsg <- response
		}
	}
}

func handleWrite(conn net.Conn) {
	encoder := json.NewEncoder(conn)

	for {
		msg := <-WriteMsg
		if err := encoder.Encode(msg); err != nil {
			writeLogPriviledged(err, "Error at handleWrite:S1")
			os.Exit(1)
		}
	}
}

func heartBeat() {
	for {
		time.Sleep(5 * time.Second)
		WriteMsg <- global.Message{Type: "ping"}
		select {
		case <-time.After(10 * time.Second):
			if global.Helper {
				writeLogPriviledged(errors.New("Patchman unexpectedly quit."), "heartBeat():S1")
				os.Exit(1)
			} else {
				global.FatalError(errors.New("Helper unexpectedly quit."))
			}
		case <-pongMsg:
			continue
		}
	}
}

func LaunchHelper() {
	global.Helper = checkAdmin()

	if global.Helper {
		initHelper()
	} else {
		initParent()
	}
}

func checkAdmin() bool {
	return os.Getuid() == 0
}

func writeLogPriviledged(err error, simplified string) {
	logfile, err1 := os.OpenFile("./patchman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("Fatal Error: Helper could not create or write to log file, simplified error: " + simplified)
		os.Exit(1)
	}
	defer logfile.Close()

	timestamp := time.Now().Truncate(time.Second).String()
	if _, err1 := logfile.WriteString(timestamp + ": " + err.Error() + "\n"); err1 != nil {
		fmt.Println("Fatal Error: Helper could not write to log file, simplified error: " + simplified)
		os.Exit(1)
	}
}
