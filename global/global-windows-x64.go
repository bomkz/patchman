//go:build windows && amd64

package global

import (
	_ "embed"
)

//go:embed patchman-unity-windows-x64.exe
var PatchmanUnity []byte
