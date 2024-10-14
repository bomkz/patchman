#Requires AutoHotkey v2.0
#include OCR.ahk

Replace(){
    rwrTexture("C:\Users\dazzl\Downloads\ahk\RWRBetty2\rwrTexture")
    SoundReplace("ttsw_shoot","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_shoot" )
    SoundReplace("ttsw_flareLow","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_flareLow" )
    SoundReplace("ttsw_pitbull","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_pitbull" )
    SoundReplace("radarLockLoop","C:\Users\dazzl\Downloads\ahk\RWRBetty2\radarLockLoop" )
    SoundReplace("ttsw_fuelDump","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_fuelDump" )
    SoundReplace("ttsw_auxEngineFailure","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_auxEngineFailure" )
    SoundReplace("ttsw_landingGear","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_landingGear" )
    SoundReplace("ttsw_pullUp","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_pullUp" )
    ReplaceEmpty("ttsw_chaffEmpty","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_chaffEmpty" )
    ReplaceMissileLockLoop()
    SoundReplace("rwrNewContact2","C:\Users\dazzl\Downloads\ahk\RWRBetty2\rwrNewContact2" )
    ReplaceNormal("ttsw_flare","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_flare" )
    SoundReplace("ttsw_chaffLow","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_chaffLow" )
    SoundReplace("mwsTone","C:\Users\dazzl\Downloads\ahk\RWRBetty2\mwsTone" )
    SoundReplace("rwrPing2","C:\Users\dazzl\Downloads\ahk\RWRBetty2\rwrPing2" )
    SoundReplace("ttsw_missileLaunch","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_missileLaunch" )
    SoundReplace("ttsw_bingoFuel","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_bingoFuel" )
    SoundReplace("ttsw_altitude","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_altitude" )
    SoundReplace("ttsw_wingFold","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_wingFold" )
    ReplaceWarningBeep()
    SoundReplace("ttsw_overG","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_overG" )
    SoundReplace("ttsw_leftEngineFailure","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_leftEngineFailure" )
    ReplaceEmpty("ttsw_flareEmpty","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_flareEmpty" )
    ReplaceNormal("ttsw_chaff","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_chaff" )
    SoundReplace("SARHLockLoop","C:\Users\dazzl\Downloads\ahk\RWRBetty2\SARHLockLoop" )
    SoundReplace("ttsw_rightEngineFailure","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_rightEngineFailure" )
    SoundReplace("ttsw_engineFailure","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_engineFailure" )
    SoundReplace("ttsw_autopilotOff","C:\Users\dazzl\Downloads\ahk\RWRBetty2\ttsw_autopilotOff" )

}


ReplaceWarningBeep(){
    Find("warningBeep")
    Sleep 500
    Send "{F3} {F3} {F3}"
    Sleep 1000
    LeftClickOnText("Data")
    Sleep 100
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 300
    SendText "C:\Users\dazzl\Downloads\ahk\RWRBetty2\warningBeep"
    Sleep 500
    ClickOpen()
    Sleep 100
    MiddleClickOnText("View Asset")
    Sleep 100
    Send "{Left}{Left}{Enter}"
    Sleep 100
}


ReplaceNormal(assetName,assetPath){

    Find(assetName)
    Sleep 500
    Send "{F3} {F3}"
    Sleep 1000
    LeftClickOnText("Data")
    Sleep 100
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 300
    SendText assetPath
    Sleep 500
    ClickOpen()
    Sleep 100
    MiddleClickOnText("View Asset")
    Sleep 100
    Send "{Left}{Left}{Enter}"
    Sleep 100

}

ReplaceEmpty(assetName, assetPath){

    Find(assetName)
    Sleep 500

    Send "{F3}"
    Sleep 1000
    LeftClickOnText("Data")
    Sleep 100
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 300
    SendText assetPath
    Sleep 500
    ClickOpen()
    Sleep 100
    MiddleClickOnText("View Asset")
    Sleep 100
    Send "{Left}{Left}{Enter}"
    Sleep 100

}

ReplaceMissileLockLoop(){

    Find("missileLockLoop")
    Sleep 500
    Send "{F3}{F3}{F3}{F3}"
    Sleep 1000
    LeftClickOnText("Data")
    Sleep 100
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 300
    SendText "C:\Users\dazzl\Downloads\ahk\RWRBetty2\missileLockLoop"
    Sleep 500
    ClickOpen()
    Sleep 100
    MiddleClickOnText("View Asset")
    Sleep 100
    Send "{Left}{Left}{Enter}"
    Sleep 100
}

NavigateToData(){
    Send "{Click 910 35}"
    SendText "C:\Program Files (x86)\Steam\steamapps\common\VTOL VR\VTOLVR_Data"
    Sleep 100
    Send "{Enter}"
}

OpenResources(){
    Send "{Click 527 181}"
    SendText "R"
    Sleep 50
    SendText "R"

    Send "{shift down} {Down} {Down} {shift up} {Enter}"

    Sleep 100 
  

    
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
    Sleep 500
    LeftClickOnText("Search by name")
    Sleep 500
    HighlightQuery()
    Sleep 300
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
    SendText "C:\Users\dazzl\Downloads\ahk\RWRBetty2"
    Sleep 100
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
    Sleep 100
    result := OCR.FromDesktop()
    try found := result.FindString("Select All")
    if !IsSet(found) {
        MsgBox 'View Data not found !'
        ExitApp
    }

    result.Click(found)
}
rwrTexture(assetPath){
   
    Find("rwrTexture")
    Sleep 500
    LeftClickOnText("Data")
    Sleep 100
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 300
    SendText assetPath
    Sleep 500
    ClickOpen()
    Sleep 1000
    MiddleClickOnText("View Asset")
    Sleep 100
    Send "{Left}{Left}{Enter}"
    Sleep 100
}

SoundReplace(assetName,assetPath){

    Find(assetName)
    Sleep 500
    LeftClickOnText("Data")
    Sleep 100
    Send "{Right}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{Down}{AppsKey}{Down}{Down}{Enter}"
    Sleep 300
    SendText assetPath
    Sleep 500
    ClickOpen()
    Sleep 100
    MiddleClickOnText("View Asset")
    Sleep 100
    Send "{Left}{Left}{Enter}"
    Sleep 100

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