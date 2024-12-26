#Requires AutoHotkey v2.0
#Include delays.ahk
#Include UIUtils.ahk

ClickOpenFileUABE(){
    gndelay
    uabeclickFile
    uidelay
    uabeClickOpenFile
}

ClickEditDataAva(){
    CoordMode "Mouse", "Screen"
    Click 1770, 400
}

OpenFileAva(){
    gndelay
    Send "{Ctrl Down}o"
    gndelay
    Send "{Ctrl Up}"
    gndelay
}

ClickExportAva(){
    CoordMode "Mouse", "Screen"
    Click 525, 150
}

ClickLZMAAva(){
    CoordMode "Mouse", "Screen"
    Click 985, 570
}

ClickAllFilesUABE(){
    gndelay
    ClickMainFileUABE
    mcdelay
    Send "{Shift Down}"
    uabeClickLastFile
    mcdelay
    Send "{Shift Up}"
    gndelay
}

ClickCloseFilesUABE(){
    gndelay
    uabeClickFile
    mcdelay
    uabeClickClose
    gndelay
}

ClickMainFileUABE(){
    CoordMode "Mouse", "Screen"
    gndelay
    Click 50, 120
    gndelay
}

ClickAssetAva(){
    gndelay
    avaClickList
    gndelay
    avaClickAsset
    gndelay
}

ClickMemoryAva(){
    CoordMode "Mouse", "Screen"
    Click 970, 570
}

ClickYesAva(){
    CoordMode "Mouse", "Screen"
    Click 900, 555
}

ClickResourceAva(){
    gndelay
    avaClickList
    gndelay
    avaClickResource
    gndelay
}

ClickImportAva(){
    CoordMode "Mouse", "Screen"
    Click 50, 210
}

ClickResSAva(){
    gndelay
    avaClickList
    gndelay
    avaClickResS
    gndelay
}

ClickFirstItemInListUABE(){
    CoordMode "Mouse", "Screen"
    Click 610, 130
}

ClickViewSearchByNameUABE(){
    gndelay 
    uabeClickView
    mcdelay
    uabeClickSearchByNameView
    gndelay
}

ClickViewItemButtonUABE(){
    CoordMode "Mouse", "Screen"
    Click 1800, 280
}
MClickItemTabUABE(){
    CoordMode "Mouse", "Screen"
    Click "middle", 800, 80
}

ClickUABE(){
    gndelay
    desktopClick
    andelay
    uabeClick
    andelay
}

GoToAvaTopItem(){
    gndelay
    ClickRandomItemAva
    gndelay
    Send "{Ctrl Down}{Home}"
    gndelay
    Send "{Ctrl up}"
}

ClickRandomItemAva(){
    CoordMode "Mouse", "Screen"
    Click 70, 500
}

ClickDataOkAva(){
    CoordMode "Mouse", "Screen"
    Click 870, 785
}

ClickAva(){
    gndelay
    desktopClick
    andelay
    avaClick()
    andelay
} 

ClickSaveAllUABE(){
    gndelay
    uabeClickFile
    mcdelay
    uabeClickApplyAndSaveAllFile
    gndelay
}