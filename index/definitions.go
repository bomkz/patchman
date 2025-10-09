package index

import "encoding/json"

var preindex PreIndexStruct

var MaxPreIndexVersion int = 0

type PreIndexStruct struct {
	Content []PreIndexContentStruct `json:"content"`
}

type PreIndexContentStruct struct {
	Version string          `json:"version"`
	Motd    string          `json:"motd,omitempty"`
	Content json.RawMessage `json:"content"`
}

var IndexURL = "https://github.com/bomkz/patchman-index/releases/latest/download/index.json"

var indexmem []byte

var useIndexVersion int
