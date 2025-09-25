package indexOne

var index []IndexStruct

type IndexStruct struct {
	Version     string                              `json:"version"`
	ObjectID    string                              `json:"id"`
	Name        string                              `json:"name"`
	Descritpion string                              `json:"description"`
	Versions    []IndexContentContentVersionsStruct `json:"versions"`
}

type IndexContentContentVersionsStruct struct {
	Version string                                     `json:"version"`
	Forms   []IndexContentContentVersionsFormsStruct   `json:"forms"`
	Content []IndexContentContentVersionsContentStruct `json:"content"`
}

type IndexContentContentVersionsContentStruct struct {
	ObjectID   string                                          `json:"id"`
	UUID       string                                          `json:"uuid"`
	Items      []IndexContentContentVersionsContentItemsStruct `json:"items"`
	PatchURL   string                                          `json:"patchURL"`
	UnpatchURL string                                          `json:"unpatchURL"`
}

type IndexContentContentVersionsContentItemsStruct struct {
	ObjectID  string `json:"id"`
	Available bool   `json:"available"`
}

type IndexContentContentVersionsFormsStruct struct {
	ObjectID string `json:"id"`
	FormType string `json:"type"`
	Name     string `json:"name"`
}

type PossibleCombinationStruct struct {
	Name      string
	ObjectIDs []string
}

type selectionstruct struct {
	Name         string
	IndexID      int
	ObjectID     string
	Description  string
	Combinations []PossibleCombinationStruct
	Combination  []string
	VariantID    string
}

var selection selectionstruct

var zstdURL = "https://github.com/bomkz/patchman-resources/releases/download/76f53ddd-7484-465e-a349-a63e35f84dc7/zstd.exe"

type StatusStruct struct {
	InstalledObjectId          string `json:"objectId,omitempty"`
	InstalledUUID              string `json:"objectUUID,omitempty"`
	InstalledVersionId         string `json:"versionId,omitempty"`
	InstalledVariantId         string `json:"variantId,omitempty"`
	InstalledVariantPatchURL   string `json:"variantPatchURL,omitempty"`
	InstalledVariantUnpatchURL string `json:"variantUnpatchURL,omitempty"`
	InstalledName              string `json:"variantName,omitempty"`
	InstalledVersion           int    `json:"installedVersion,omitempty"`
}
type StatusTargetStruct struct {
	TargetObjectId          string
	TargetUUID              string
	TargetVersionId         string
	TargetVariantId         string
	TargetVariantPatchURL   string
	TargetVariantUnpatchURL string
}

var StatusTarget StatusTargetStruct

var Status StatusStruct
