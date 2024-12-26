#Requires AutoHotkey v2.0
#Include utils/resources.ahk

EFResourceFile  := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061"
EFArchiveName   := "CAB-db515831ae078197daa2fd6af388d061/"

OpenFileInAva(EFResourceFile)
FixEFResources
SaveModifiedResourcesAva

FixEFResources(){
    FixTextures("dashSprites",          "",     EFArchiveName)
    FixTextures("rwrTexture",           "{F3}", EFArchiveName)
    FixAudio("missileLockLoop",         "",     EFArchiveName)
    FixAudio("mwsTone",                 "",     EFArchiveName)
    FixAudio("radarLockLoop",           "",     EFArchiveName)
    FixAudio("rwrNewContact2",          "",     EFArchiveName)
    FixAudio("rwrPing2",                "",     EFArchiveName)
    FixAudio("SARHLockLoop",            "",     EFArchiveName)
    FixAudio("ttsw_altitude",           "",     EFArchiveName)
    FixAudio("ttsw_autopilotOff",       "",     EFArchiveName)
    FixAudio("ttsw_auxEngineFailure",   "",     EFArchiveName)
    FixAudio("ttsw_bingoFuel",          "",     EFArchiveName)
    FixAudio("ttsw_chaff",              "",     EFArchiveName)
    FixAudio("ttsw_chaffEmpty",         "",     EFArchiveName)
    FixAudio("ttsw_chaffLow",           "",     EFArchiveName)
    FixAudio("ttsw_flare",              "",     EFArchiveName)
    FixAudio("ttsw_flareEmpty",         "",     EFArchiveName)
    FixAudio("ttsw_flareLow",           "",     EFArchiveName)
    FixAudio("ttsw_fuelDump",           "",     EFArchiveName)
    FixAudio("ttsw_landingGear",        "",     EFArchiveName)
    FixAudio("ttsw_leftEngineFailure",  "",     EFArchiveName)
    FixAudio("ttsw_missileLaunch",      "",     EFArchiveName)
    FixAudio("ttsw_overG",              "",     EFArchiveName)
    FixAudio("ttsw_pitbull",            "",     EFArchiveName)
    FixAudio("ttsw_pullUp",             "",     EFArchiveName)
    FixAudio("ttsw_rightEngineFailure", "",     EFArchiveName)
    FixAudio("ttsw_shoot",              "",     EFArchiveName)
    FixAudio("ttsw_wingFold",           "",     EFArchiveName)
    FixAudio("warningBeep",             "",     EFArchiveName)
}