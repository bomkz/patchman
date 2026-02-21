package procman

import (
	"github.com/bomkz/patchman/ipc/procman/platforms"
)

func InitProcess() {
	platforms.LaunchHelper()
}
