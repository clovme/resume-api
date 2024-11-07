echo off
go build -ldflags "-s -w" -trimpath -o resume.exe
upx --ultra-brute --best --lzma --brute --compress-exports=1 --no-mode --no-owner --no-time --force resume.exe
resume.exe