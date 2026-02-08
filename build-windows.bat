@echo off
echo Building NetFlow for Windows...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w -H windowsgui" -o netflow-windows-amd64.exe
if %ERRORLEVEL% EQU 0 (
    echo Build successful: netflow-windows-amd64.exe
) else (
    echo Build failed!
    exit /b 1
)

