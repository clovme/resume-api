@echo off
go build -ldflags "-s -w" -trimpath -v -x -o "resume.exe"
upx --ultra-brute --best --lzma --brute --compress-exports=1 --no-mode --no-owner --no-time --force "resume.exe"