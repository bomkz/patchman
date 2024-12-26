Start-Sleep 15

$ErrorActionPreference = "Stop"

steamcmd.exe +force_install_dir "C:\Users\Public\Desktop\vtolvr\" +login st0306 +app_update 667970 +quit

Invoke-WebRequest -o C:\betty.zip https://codeload.github.com/bomkz/rwrbettyassets/zip/refs/tags/betty2
Expand-Archive C:\betty.zip C:\ 
Move-Item C:\rwrbettyassets-betty2\resources C:\
Remove-Item C:\rwrbettyassets-betty2
Remove-Item C:\betty.zip

Start-Process -FilePath C:\Users\Public\Desktop\scripts\basemodbuilder.ahk -Wait 
Start-Process -FilePath C:\Users\Public\Desktop\scripts\ef24modbuilder.ahk -Wait
Start-Process -FilePath C:\Users\Public\Desktop\scripts\ah94modbuilder.ahk -Wait 

Remove-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea
Remove-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS
Remove-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource

Move-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea-mod C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea
Move-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS-mod C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS
Move-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource-mod C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource

Remove-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061
Remove-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS
Remove-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource

Move-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061-mod C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061
Move-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS-mod C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS
Move-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource-mod C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource

Start-Process -FilePath C:\Users\Public\Desktop\scripts\ef24modfixer.ahk -Wait
Start-Process -FilePath C:\Users\Public\Desktop\scripts\ah94modfixer.ahk -Wait

Start-Process -FilePath C:\Users\Public\Desktop\scripts\ah94modpacker.ahk -Wait
Start-Process -FilePath C:\Users\Public\Desktop\scripts\ef24modpacker.ahk -Wait

Remove-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea
Remove-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS
Remove-Item C:\Users\Public\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource

Remove-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061
Remove-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS
Remove-Item C:\Users\Public\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource

Remove-Item C:\Users\Public\Desktop\1770480-tmp
Remove-Item C:\Users\Public\Desktop\2531290-tmp

Remove-Item C:\resources

C:\Users\Public\Desktop\zstd.exe --patch-from="C:\Users\Public\Desktop\vtolvr\VTOLVR_Data\resources.resource" "C:\Users\Public\Desktop\resources.resource" -o "C:\Users\Public\Desktop\resources.resource.patch"
C:\Users\Public\Desktop\zstd.exe --patch-from="C:\Users\Public\Desktop\vtolvr\VTOLVR_Data\resources.assets" "C:\Users\Public\Desktop\resources.assets" -o "C:\Users\Public\Desktop\resources.assets.patch"
C:\Users\Public\Desktop\zstd.exe --patch-from="C:\Users\Public\Desktop\vtolvr\VTOLVR_Data\resources.assets.resS" "C:\Users\Public\Desktop\resources.assets.resS" -o "C:\Users\Public\Desktop\resources.assets.resS.patch"

Remove-Item  C:\Users\Public\Desktop\resources.assets
Remove-Item  C:\Users\Public\Desktop\resources.assets.resS
Remove-Item  C:\Users\Public\Desktop\resources.resource

$compress = @{
    Path = "C:\Users\Public\Desktop\resources.assets.patch", "C:\Users\Public\Desktop\resources.assets.resS.patch", "C:\Users\Public\Desktop\resources.resource.patch", "C:\Users\Public\Desktop\zstd.exe"
    CompressionLevel = "Optimal"
    DestinationPath = "C:\Archives\f16-mod.zip"
}

Compress-Archive @compress

[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5
[Console]::Beep()
Start-Sleep 5