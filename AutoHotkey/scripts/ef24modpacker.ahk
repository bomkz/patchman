#Requires AutoHotkey v2.0
#Include utils/resources.ahk

EFBundleTmpPath     := "C:\Users\Public\Desktop\2531290-tmp"
EFBundleExportPath  := "C:\Users\Public\Desktop\2531290"
EFBundlePath        := "C:\Users\Public\Desktop\vtolvr\DLC\2531290\2531290"
EFResourceFile1     := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061"
EFResourceFile2     := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource"
EFResourceFile3     := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS"

OpenBundle(EFBundlePath)
ReImportAvaFiles(EFResourceFile3,EFResourceFile1,EFResourceFile2)
SaveFileAva(EFBundleTmpPath)
CleanupAva
OpenBundle(EFBundleTmpPath)
CompressFileAva(EFBundleExportPath)
CleanupAva