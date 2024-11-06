#Requires AutoHotkey v2.0
#include OCR.ahk

Replace(){
    dashTexture("C:\ahk\RWRBetty2\dashSprites")
    rwrTexture("C:\ahk\RWRBetty2\rwrTexture")
    SoundReplace("ttsw_shoot","C:\ahk\RWRBetty2\ttsw_shoot" )
    SoundReplace("ttsw_flareLow","C:\ahk\RWRBetty2\ttsw_flareLow" )
    SoundReplace("ttsw_pitbull","C:\ahk\RWRBetty2\ttsw_pitbull" )
    SoundReplace("radarLockLoop","C:\ahk\RWRBetty2\radarLockLoop" )
    SoundReplace("ttsw_fuelDump","C:\ahk\RWRBetty2\ttsw_fuelDump" )
    SoundReplace("ttsw_auxEngineFailure","C:\ahk\RWRBetty2\ttsw_auxEngineFailure" )
    SoundReplace("ttsw_landingGear","C:\ahk\RWRBetty2\ttsw_landingGear" )
    SoundReplace("ttsw_pullUp","C:\ahk\RWRBetty2\ttsw_pullUp" )
    ReplaceEmpty("ttsw_chaffEmpty","C:\ahk\RWRBetty2\ttsw_chaffEmpty" )
    ReplaceMissileLockLoop()
    SoundReplace("rwrNewContact2","C:\ahk\RWRBetty2\rwrNewContact2" )
    ReplaceNormal("ttsw_flare","C:\ahk\RWRBetty2\ttsw_flare" )
    SoundReplace("ttsw_chaffLow","C:\ahk\RWRBetty2\ttsw_chaffLow" )
    SoundReplace("mwsTone","C:\ahk\RWRBetty2\mwsTone" )
    SoundReplace("rwrPing2","C:\ahk\RWRBetty2\rwrPing2" )
    SoundReplace("ttsw_missileLaunch","C:\ahk\RWRBetty2\ttsw_missileLaunch" )
    SoundReplace("ttsw_bingoFuel","C:\ahk\RWRBetty2\ttsw_bingoFuel" )
    SoundReplace("ttsw_altitude","C:\ahk\RWRBetty2\ttsw_altitude" )
    SoundReplace("ttsw_wingFold","C:\ahk\RWRBetty2\ttsw_wingFold" )
    ReplaceWarningBeep()
    SoundReplace("ttsw_overG","C:\ahk\RWRBetty2\ttsw_overG" )
    SoundReplace("ttsw_leftEngineFailure","C:\ahk\RWRBetty2\ttsw_leftEngineFailure" )
    ReplaceEmpty("ttsw_flareEmpty","C:\ahk\RWRBetty2\ttsw_flareEmpty" )
    ReplaceNormal("ttsw_chaff","C:\ahk\RWRBetty2\ttsw_chaff" )
    SoundReplace("SARHLockLoop","C:\ahk\RWRBetty2\SARHLockLoop" )
    SoundReplace("ttsw_rightEngineFailure","C:\ahk\RWRBetty2\ttsw_rightEngineFailure" )
    SoundReplace("ttsw_engineFailure","C:\ahk\RWRBetty2\ttsw_engineFailure" )
    SoundReplace("ttsw_autopilotOff","C:\ahk\RWRBetty2\ttsw_autopilotOff" )

}


ReplaceWarningBeep(){
    Find("warningBeep")
    Sleep 600
    Send "{F3} {F3} {F3}"
    Sleep 1100
    LeftClickOnText("Data")
    Sleep 200
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 400
    SendText "C:\ahk\RWRBetty2\warningBeep"
    Sleep 600
    ClickOpen()
    Sleep 200
    MiddleClickOnText("View Asset")
    Sleep 200
    Send "{Left}{Left}{Enter}"
    Sleep 200
}


ReplaceNormal(assetName,assetPath){

    Find(assetName)
    Sleep 600
    Send "{F3} {F3}"
    Sleep 1100
    LeftClickOnText("Data")
    Sleep 200
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 400
    SendText assetPath
    Sleep 600
    ClickOpen()
    Sleep 200
    MiddleClickOnText("View Asset")
    Sleep 200
    Send "{Left}{Left}{Enter}"
    Sleep 200

}

ReplaceEmpty(assetName, assetPath){

    Find(assetName)
    Sleep 600
    Send "{F3}"
    Sleep 1100
    LeftClickOnText("Data")
    Sleep 200
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 400
    SendText assetPath
    Sleep 600
    ClickOpen()
    Sleep 200
    MiddleClickOnText("View Asset")
    Sleep 200
    Send "{Left}{Left}{Enter}"
    Sleep 200

}

ReplaceMissileLockLoop(){

    Find("missileLockLoop")
    Sleep 600
    Send "{F3}{F3}{F3}{F3}"
    Sleep 1100
    LeftClickOnText("Data")
    Sleep 200
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 400
    SendText "C:\ahk\RWRBetty2\missileLockLoop"
    Sleep 600
    ClickOpen()
    Sleep 200
    MiddleClickOnText("View Asset")
    Sleep 200
    Send "{Left}{Left}{Enter}"
    Sleep 200
}

NavigateToData(){
    Send "{Click 910 35}"
    SendText "C:\Program Files (x86)\Steam\steamapps\common\VTOL VR\VTOLVR_Data"
    Sleep 200
    Send "{Enter}"
}

OpenResources(){
    Send "{Click 527 181}"
    SendText "R"
    Sleep 150
    SendText "R"

    Send "{shift down} {Down} {Down} {shift up} {Enter}"

    Sleep 200 
  

    
}

ClickSearchByName(){
    result := OCR.FromDesktop()
    try found := result.FindString("Search by name")
    if !IsSet(found) {
        MsgBox '"Search by name" was not found in UABE !'
        ExitApp
    }

    result.Click(found)
}

Find(assetName){
    Send "{Click 640 500} {Home}"
    LeftClickOnText("View")
    Sleep 600
    LeftClickOnText("Search by name")
    Sleep 600
    HighlightQuery()
    Sleep 400
    SendText assetName
    Send "{Enter}"
}



ClickOpen(){
    result := OCR.FromDesktop()
    try found := result.FindString("Any file")
    if !IsSet(found) {
        MsgBox 'not found Any file !' 
        ExitApp
    }

    CoordMode "Mouse", "Screen"
    Click found.x, found.y+40 
}



MiddleClickOnText(inputString){
    result := OCR.FromDesktop()
    try found := result.FindString(inputString)
    if !IsSet(found) {
        MsgBox 'not found !' inputString
        ExitApp
    }

    result.Click(found, "Middle")
}

LeftClickOnText(inputString){
    result := OCR.FromDesktop()
    try found := result.FindString(inputString)
    if !IsSet(found) {
        MsgBox 'not found !' inputString
        ExitApp
    }

    result.Click(found)
}
RightClickOnText(inputString){

    result := OCR.FromDesktop()
    try found := result.FindString(inputString)
    if !IsSet(found) {
        MsgBox 'not found !' inputString
        ExitApp
    }

   

    result.Click(found, "Right")
}



NavigateToAssets(){
    Send "{Click 910 35}"
    SendText "C:\ahk\RWRBetty2"
    Sleep 200
    Send "{Enter}"
}


HighlightQuery(){
    result := OCR.FromDesktop()
    try found := result.FindString("(* allowed)")
    if !IsSet(found) {
        MsgBox 'not found Any file !' 
        ExitApp
    }

    CoordMode "Mouse", "Screen"
    Click found.x+100, found.y, "Right"
    Sleep 200
    result := OCR.FromDesktop()
    try found := result.FindString("Select All")
    if !IsSet(found) {
        MsgBox 'View Data not found !'
        ExitApp
    }

    result.Click(found)
}

dashTexture(assetPath){
   
    Find("dashSprites")
    Sleep 600
    LeftClickOnText("Data")
    Sleep 200
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 400
    SendText assetPath
    Sleep 600
    ClickOpen()
    Sleep 1100
    MiddleClickOnText("View Asset")
    Sleep 200
    Send "{Left}{Left}{Enter}"
    Sleep 200
}

rwrTexture(assetPath){
   
    Find("rwrTexture")
    Sleep 600
    LeftClickOnText("Data")
    Sleep 200
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 400
    SendText assetPath
    Sleep 600
    ClickOpen()
    Sleep 1100
    MiddleClickOnText("View Asset")
    Sleep 200
    Send "{Left}{Left}{Enter}"
    Sleep 200
}

SoundReplace(assetName,assetPath){

    Find(assetName)
    Sleep 600
    LeftClickOnText("Data")
    Sleep 200
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 400
    SendText assetPath
    Sleep 600
    ClickOpen()
    Sleep 200
    MiddleClickOnText("View Asset")
    Sleep 200
    Send "{Left}{Left}{Enter}"
    Sleep 200

}

ClickName(){
    result := OCR.FromDesktop()
    try found := result.FindString("Size (bytes)")
    if !IsSet(found) {
        MsgBox 'not found meow!' 
        ExitApp
    }

    CoordMode "Mouse", "Screen"
    Click found.x-400, found.y 
}