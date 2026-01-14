package indexHandler

import "encoding/json"

var preindex PreIndexStruct

// Index version patchman should use
var PreIndexVersion int = 2

// Index root
type PreIndexStruct struct {
	Content []PreIndexContentStruct `json:"content"`
}

// Index struct used for index array in PreIndexArray
type PreIndexContentStruct struct {
	Version string          `json:"version"`
	Motd    string          `json:"motd,omitempty"`
	Content json.RawMessage `json:"content"`
}

// Url for Index
var IndexURL = "https://github.com/bomkz/patchman-index/releases/latest/download/index.json"

// Stores Index Json
var indexmem []byte

var useIndexVersion int
