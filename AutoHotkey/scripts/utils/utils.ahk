#Requires AutoHotkey v2.0

#Include UI.ahk
#Include delays.ahk

OpenUABEFile(fileName) {
    gndelay
    ClickOpenFileUABE
    uidelay
    SendText filename
    tidelay
    Send "{Enter}"
    gndelay
}
ExportAvaFile(bundleFilePath,exportPath) {
    gndelay
    OpenFileAva
    andelay
    SendText bundleFilePath
    tidelay
    Send "{Enter}"
    andelay
    ClickMemoryAva
    dcdelay
    ClickExportAva 
    andelay
    SendText exportPath
    tidelay
    Send "{Enter}"
    tidelay
    Send "{Enter}"
    andelay
}

ViewItemData(itemName, extraKeystrokes) {
    gndelay
    SearchItem(itemName)
    gndelay
    Send extraKeystrokes
    tidelay
    ClickViewItemButtonUABE
}



SaveItem(){
    gndelay
    MClickItemTabUABE
    uidelay
    YesNoCancelSelectYes        
    gndelay
}

YesNoCancelSelectYes(){
    gndelay
    Send "{Left}"
    smdelay
    Send "{Left}"
    smdelay
    Send "{Enter}"
    gndelay
}


ImportResource(itemFilePath){
    gndelay
    SendText itemFilePath
    tidelay
    Send "{Enter}"
    gndelay
}

OpenImportMenu(){
    gndelay
    Send "{AppsKey}"
    gndelay
    Send "{Down}"
    gndelay
    Send "{Down}"
    gndelay
    Send "{Enter}"
    uidelay
}


HighlightAudioResource(){
    Send "{Right}"
    smdelay
    Send "{PgDn}"
    smdelay
    Send "{Up}"
    gndelay    
}


SearchItem(itemName){ 
    gndelay
    ClickViewSearchByNameUABE
    uidelay
    Send "{Tab}"
    gndelay
    Send "{Tab}"
    gndelay
    SendText(itemName)
    tidelay
    Send "{Enter}"
    gndelay
}

HighlightTextureResource(){
    gndelay
    Send "{Right}"
    smdelay
    Send "{PgDn}"
    gndelay
}

returnToTop(){
    gndelay
    ClickFirstItemInListUABE
    gndelay
    Send "{Home}"
    gndelay
}
