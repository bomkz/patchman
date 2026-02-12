//go:build darwin

package global

import (
	_ "embed"
)

//go:embed patchman-unity-darwin-aio
var PatchmanUnity []byte
