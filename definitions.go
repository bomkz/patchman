package main

//go:embed installer.zip
var installerfiles []byte
var resourcesresourcepatch []byte
var resourcesassetspatch []byte
var resourcesassetsresspatch []byte
var ah94 []byte
var ef24g []byte
var AH94Installed = false
var EF24GInstalled = false
var vtolvrpath string

type SteamLibraryFolder struct {
	Path string `json:"path,omitempty"`
}
