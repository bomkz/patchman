//go:build windows

package platforms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/Microsoft/go-winio"
	"github.com/bomkz/patchman/global"
	"golang.org/x/sys/windows"
)

var (
	shellExecuteEx = windows.MustLoadDLL("shell32.dll").MustFindProc("ShellExecuteExW")
	WriteMsg       = make(chan global.Message)
	ReadMsg        = make(chan global.Message)
)

const (
	pipeName                = `\\.\pipe\patchman-pipe`
	SEE_MASK_NOCLOSEPROCESS = 0x00000040
	SW_HIDE                 = 0
)

func initParent() {
	parentPID := os.Getpid()

	l := global.Assure(winio.ListenPipe(pipeName, &winio.PipeConfig{
		MessageMode: false, // byte mode works fine with JSON
	}))

	childPID := launchHelper()
	allowedPIDs := map[int]bool{
		parentPID: true,
		childPID:  true,
	}

	conn := acceptAndVerify(l, allowedPIDs)

	go handleRead(conn)
	go handleWrite(conn)
	//go heartBeat()

}

func initHelper() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var conn net.Conn
	for {
		var err error
		conn, err = winio.DialPipeContext(ctx, pipeName)
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
	//go heartBeat()
}

func launchHelper() int {
	exe, _ := os.Executable()
	exe, _ = filepath.EvalSymlinks(exe)

	verb, _ := windows.UTF16PtrFromString("runas")
	file, _ := windows.UTF16PtrFromString(exe)
	type SHELLEXECUTEINFO struct {
		cbSize         uint32
		fMask          uint32
		hwnd           uintptr
		lpVerb         *uint16
		lpFile         *uint16
		lpParameters   *uint16
		lpDirectory    *uint16
		nShow          int32
		hInstApp       uintptr
		lpIDList       uintptr
		lpClass        *uint16
		hkeyClass      uintptr
		dwHotKey       uint32
		hIconOrMonitor uintptr
		hProcess       windows.Handle
	}

	info := &SHELLEXECUTEINFO{
		fMask:  SEE_MASK_NOCLOSEPROCESS,
		lpVerb: verb,
		lpFile: file,
		nShow:  SW_HIDE,
	}
	info.cbSize = uint32(unsafe.Sizeof(*info))

	ret, _, err := shellExecuteEx.Call(uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		global.FatalError(err)
	}

	// hProcess is valid because of SEE_MASK_NOCLOSEPROCESS
	pid := global.Assure(windows.GetProcessId(info.hProcess))

	windows.CloseHandle(info.hProcess)

	return int(pid)
}

func acceptAndVerify(l net.Listener, allowedPIDs map[int]bool) net.Conn {
	conn, err := l.Accept()
	if err != nil {
		global.FatalError(err)
		return nil
	}

	// winio pipe connections expose the underlying handle via the
	// internal file — use GetNamedPipeClientProcessId on the raw handle
	type hasFD interface {
		Fd() uintptr
	}
	// winio returns a *net.pipeConn; cast through the file descriptor
	pipeConn, ok := conn.(hasFD)
	if !ok {
		conn.Close()
		global.FatalError(fmt.Errorf("connection does not expose file descriptor"))
		return nil
	}

	var pid uint32
	handle := windows.Handle(pipeConn.Fd())
	global.AssureNoReturn(windows.GetNamedPipeClientProcessId(handle, &pid))

	if !allowedPIDs[int(pid)] {
		conn.Close()
		global.FatalError(fmt.Errorf("unauthorized PID: %d", pid))
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
			os.Exit(1) // covers EOF (other end closed) and any read error
		}
		switch response.Type {
		case "ping":
			pong := global.Message{Type: "pong"}
			WriteMsg <- pong
			continue
		case "pong":
			pongMsg <- true
			continue
		default:
			ReadMsg <- response

		}
	}
}

func heartBeat() {
	for {
		time.Sleep(1000 * time.Second)
		var ping = global.Message{Type: "ping"}
		WriteMsg <- ping
		select {
		case <-time.Tick(10000 * time.Second):
			if global.Helper {
				// Section 1
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

var pongMsg = make(chan bool)

func handleWrite(conn net.Conn) {
	encoder := json.NewEncoder(conn)

	for {

		msg := <-WriteMsg

		if err := encoder.Encode(msg); err != nil {
			// Section 1
			writeLogPriviledged(err, "Error at handleWrite:S1")
			os.Exit(1)
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

// Returns whether running as admin or not.
func checkAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	isadmin := false
	if err != nil {
		isadmin = false
	} else {
		isadmin = true
	}
	return isadmin
}

func writeLogPriviledged(err error, simplified string) {

	logfile, err1 := os.OpenFile("./patchman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("Fatal Error: Helper could not create or write to log file, simplified error: " + simplified)
		fmt.Println("exit")
		os.Exit(1)
	}
	timestamp := time.Now().Truncate(time.Second).String()

	if _, err1 := logfile.WriteString(timestamp + ": " + err.Error() + "\n"); err1 != nil {
		fmt.Println("Fatal Error: Helper could not create or write to log file, simplified error: " + simplified)
		fmt.Println("exit")
		logfile.Close()
		os.Exit(1)
	}

	logfile.Close()

}
