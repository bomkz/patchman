//go:build windows && arm64

package global

import (
	_ "embed"
)

//go:embed patchman-unity-windows-arm64.exe
var PatchmanUnity []byte
