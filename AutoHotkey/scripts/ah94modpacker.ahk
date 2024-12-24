#Requires AutoHotkey v2.0
#Include utils/resources.ahk

BundleTmpPath       := "C:\Users\Administrator\Desktop\1770480-tmp"
BundleExportPath    := "C:\Users\Administrator\Desktop\1770480"
BundlePath          := "C:\Users\Administrator\Desktop\vtolvr\DLC\1770480\1770480"
ResourceFile1       := "C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea"
ResourceFile2       := "C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource"
ResourceFile3       := "C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS"
AssetsPath          := "C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea-mod"
ResourcePath        := "C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource-mod"
ResSPath            := "C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS-mod"
ExportPath          := "C:\Users\Administrator\Desktop\"



OpenBundle(BundlePath)
ReImportAvaFiles(ResourceFile3,ResourceFile1,ResourceFile2)
SaveFileAva(BundleTmpPath)
CleanupAva
OpenBundle(BundleTmpPath)
CompressFileAva(BundleExportPath)
CleanupAva