//go:build linux && amd64

package global

import (
	_ "embed"
)

//go:embed patchman-unity-linux-x64
var PatchmanUnity []byte
