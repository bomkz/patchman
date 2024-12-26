#Requires AutoHotkey v2.0
#Include utils/resources.ahk

AHResourceFile  := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea"
AHArchiveName   := "CAB-609a7bd01976702a18d81971aebebeea/"

OpenFileInAva(AHResourceFile)
FixAHResources
SaveModifiedResourcesAva

FixAHResources(){
    FixTextures("dashSprites",          "",     AHArchiveName)
    FixTextures("rwrTexture",           "{F3}", AHArchiveName)
    FixAudio("missileLockLoop",         "",     AHArchiveName)
    FixAudio("mwsTone",                 "",     AHArchiveName)
    FixAudio("radarLockLoop",           "",     AHArchiveName)
    FixAudio("rwrNewContact2",          "",     AHArchiveName)
    FixAudio("rwrPing2",                "",     AHArchiveName)
    FixAudio("SARHLockLoop",            "",     AHArchiveName)
    FixAudio("ttsw_altitude",           "",     AHArchiveName)
    FixAudio("ttsw_autopilotOff",       "",     AHArchiveName)
    FixAudio("ttsw_auxEngineFailure",   "",     AHArchiveName)
    FixAudio("ttsw_bingoFuel",          "",     AHArchiveName)
    FixAudio("ttsw_chaff",              "",     AHArchiveName)
    FixAudio("ttsw_chaffEmpty",         "",     AHArchiveName)
    FixAudio("ttsw_chaffLow",           "",     AHArchiveName)
    FixAudio("ttsw_flare",              "",     AHArchiveName)
    FixAudio("ttsw_flareEmpty",         "",     AHArchiveName)
    FixAudio("ttsw_flareLow",           "",     AHArchiveName)
    FixAudio("ttsw_fuelDump",           "",     AHArchiveName)
    FixAudio("ttsw_landingGear",        "",     AHArchiveName)
    FixAudio("ttsw_leftEngineFailure",  "",     AHArchiveName)
    FixAudio("ttsw_missileLaunch",      "",     AHArchiveName)
    FixAudio("ttsw_overG",              "",     AHArchiveName)
    FixAudio("ttsw_pitbull",            "",     AHArchiveName)
    FixAudio("ttsw_pullUp",             "",     AHArchiveName)
    FixAudio("ttsw_rightEngineFailure", "",     AHArchiveName)
    FixAudio("ttsw_shoot",              "",     AHArchiveName)
    FixAudio("ttsw_wingFold",           "",     AHArchiveName)
    FixAudio("warningBeep",             "",     AHArchiveName)
}