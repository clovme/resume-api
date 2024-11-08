echo off
go build -ldflags "-s -w" -trimpath -o "D:\Application\Rolan\Tools\Common Software\resume\resume.exe"
upx --ultra-brute --best --lzma --brute --compress-exports=1 --no-mode --no-owner --no-time --force "D:\Application\Rolan\Tools\Common Software\resume\resume.exe"