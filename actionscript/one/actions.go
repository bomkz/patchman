package actionScriptOne

func HandleActions(actionData []byte, actionType string) {
	batchAssetImport()
}
func batchAssetImport() {
	createOperationsFile()
	runPatchmanUnity()

}
