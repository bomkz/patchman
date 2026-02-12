//go:build linux && arm64

package global

import (
	_ "embed"
)

//go:embed patchman-unity-linux-arm64
var PatchmanUnity []byte
