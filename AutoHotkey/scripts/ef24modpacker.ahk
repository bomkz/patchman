#Requires AutoHotkey v2.0
#Include utils/resources.ahk

BundleTmpPath       := "C:\Users\Administrator\Desktop\2531290-tmp"
BundleExportPath    := "C:\Users\Administrator\Desktop\2531290"
BundlePath          := "C:\Users\Administrator\Desktop\vtolvr\DLC\2531290\2531290"
ResourceFile1       := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061"
ResourceFile2       := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource"
ResourceFile3       := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS"
AssetsPath          := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061-mod"
ResourcePath        := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource-mod"
ResSPath            := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061-mod"
ExportPath          := "C:\Users\Administrator\Desktop\"



OpenBundle(BundlePath)
ReImportAvaFiles(ResourceFile3,ResourceFile1,ResourceFile2)
SaveFileAva(BundleTmpPath)
CleanupAva
OpenBundle(BundleTmpPath)
CompressFileAva(BundleExportPath)
CleanupAva