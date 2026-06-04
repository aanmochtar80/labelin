@echo off
echo Building Labelin for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o labelin_windows_amd64.exe cmd/server/main.go
echo Build complete.
