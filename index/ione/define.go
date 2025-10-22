package ione

type IndexStruct struct {
	Version string               `json:"version"`
	Content []IndexContentStruct `json:"content"`
	MOTD    string
}

type IndexContentStruct struct {
	PatchName     string                            `json:"patchName"`
	PatchDesc     string                            `json:"patchDesc"`
	PatchAuthor   string                            `json:"patchAuthor"`
	PatchLink     string                            `json:"patchLink"`
	PatchVariants []IndexContentPatchVariantsStruct `json:"patchVariants"`
}

type IndexContentPatchVariantsStruct struct {
	Variant      string `json:"variant"`
	DownloadLink string `json:"downloadLink"`
}

var Index IndexStruct
