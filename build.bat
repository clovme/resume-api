echo off
go build -ldflags "-s -w" -trimpath -o resume.exe
D:\Application\upx\upx --best --lzma --force resume.exe
resume.exe