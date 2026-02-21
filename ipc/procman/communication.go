package procman

import (
	"github.com/bomkz/patchman/global"
	"github.com/bomkz/patchman/ipc/procman/platforms"
)

func Write(requestType string, payload []byte) {
	meowmsg := global.Message{Type: requestType, Payload: payload}
	platforms.WriteMsg <- meowmsg
}

func Read() global.Message {
	return <-platforms.ReadMsg
}
