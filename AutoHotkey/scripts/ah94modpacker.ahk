#Requires AutoHotkey v2.0
#Include utils/resources.ahk

AHPBundleTmpPath       := "C:\Users\Public\Desktop\1770480-tmp"
AHPBundleExportPath    := "C:\Users\Public\Desktop\1770480"
AHPBundlePath          := "C:\Users\Public\Desktop\vtolvr\DLC\1770480\1770480"
AHPResourceFile1       := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea"
AHPResourceFile2       := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource"
AHPResourceFile3       := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS"
AHPAssetsPath          := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea-mod"
AHPResourcePath        := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource-mod"
AHPResSPath            := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS-mod"
AHPExportPath          := "C:\Users\Public\Desktop\"



OpenBundle(AHPBundlePath)
ReImportAvaFiles(AHPResourceFile3,AHPResourceFile1,AHPResourceFile2)
SaveFileAva(AHPBundleTmpPath)
CleanupAva
OpenBundle(AHPBundleTmpPath)
CompressFileAva(AHPBundleExportPath)
CleanupAva