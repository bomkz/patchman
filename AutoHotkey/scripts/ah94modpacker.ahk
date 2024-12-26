#Requires AutoHotkey v2.0
#Include utils/resources.ahk

AHBundleTmpPath     := "C:\Users\Public\Desktop\1770480-tmp"
AHBundleExportPath  := "C:\Users\Public\Desktop\1770480"
AHBundlePath        := "C:\Users\Public\Desktop\vtolvr\DLC\1770480\1770480"
AHResourceFile1     := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea"
AHResourceFile2     := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource"
AHResourceFile3     := "C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS"

OpenBundle(AHBundlePath)
ReImportAvaFiles(AHResourceFile3,AHResourceFile1,AHResourceFile2)
SaveFileAva(AHBundleTmpPath)
CleanupAva
OpenBundle(AHBundleTmpPath)
CompressFileAva(AHBundleExportPath)
CleanupAva