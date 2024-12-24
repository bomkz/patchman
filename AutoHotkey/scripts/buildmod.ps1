Start-Sleep 15

steamcmd.exe +force_install_dir "C:\Users\Administrator\Desktop\vtolvr\" +login st0306 +app_update 667970 +quit

Invoke-WebRequest -o C:\betty.zip https://codeload.github.com/bomkz/rwrbettyassets/zip/refs/tags/betty2
Expand-Archive C:\betty.zip C:\ 
Move-Item C:\rwrbettyassets-betty2\resources C:\
Remove-Item C:\rwrbettyassets-betty2

Start-Process -FilePath C:\Users\Administrator\Desktop\scripts\basemodbuilder.ahk -Wait 
Start-Process -FilePath C:\Users\Administrator\Desktop\scripts\ef24modbuildar.ahk -Wait
Start-Process -FilePath C:\Users\Administrator\Desktop\scripts\ah94modbuilder.ahk -Wait 

Remove-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea
Remove-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS
Remove-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource

Move-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea-mod C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea
Move-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS-mod C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS
Move-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource-mode C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource

Remove-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061
Remove-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS
Remove-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource

Move-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061-mod C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061
Move-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS-mod C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS
Move-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource-mod C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource

Start-Process -FilePath C:\Users\Administrator\Desktop\scripts\ah94modbuilder.ahk -Wait
Start-Process -FilePath C:\Users\Administrator\Desktop\scripts\ef24modbuilder.ahk -Wait

Remove-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea
Remove-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resS
Remove-Item C:\Users\Administrator\Desktop\CAB-609a7bd01976702a18d81971aebebeea.resource

Remove-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061
Remove-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS
Remove-Item C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource

Remove-Item C:\Users\Administrator\Desktop\1770480-tmp
Remove-Item C:\Users\Administrator\Desktop\2531290-tmp