UHHH
Idk what to put here
Hi, this is a mod installer for my VTOL VR Mods, mostly sound asset replacements and some texture replacements.

For self reference on how to set up for compilation:

Modify versioning on: versioninfo.json modinstaller.exe.manifest 

go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go generate
go build

signtool sign /tr http://timestamp.digicert.com /td sha256 /fd sha256 /f <certificate> "modinstaller.exe"