#Requires AutoHotkey v2.0
#Include utils/resources.ahk

BundleTmpPath       := "C:\Users\Public\Desktop\2531290-tmp"
BundleExportPath    := "C:\Users\Public\Desktop\2531290"
BundlePath          := "C:\Users\Public\Desktop\vtolvr\DLC\2531290\2531290"
ResourceFile1       := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061"
ResourceFile2       := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource"
ResourceFile3       := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS"
AssetsPath          := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061-mod"
ResourcePath        := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource-mod"
ResSPath            := "C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061-mod"
ExportPath          := "C:\Users\Public\Desktop\"



OpenBundle(BundlePath)
ReImportAvaFiles(ResourceFile3,ResourceFile1,ResourceFile2)
SaveFileAva(BundleTmpPath)
CleanupAva
OpenBundle(BundleTmpPath)
CompressFileAva(BundleExportPath)
CleanupAva