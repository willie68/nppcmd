@echo off
echo building tool
go build -ldflags="-s -w" -o nppcmd.exe cmd/main.go
