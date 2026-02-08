@echo off
echo Building NetFlow for Windows...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w -H windowsgui" -o netflow-windows-amd64.exe
if %ERRORLEVEL% EQU 0 (
    echo Build successful: netflow-windows-amd64.exe
    if not exist release mkdir release
    copy /Y netflow-windows-amd64.exe release\netflow.exe >nul
    echo Copied to release\netflow.exe for README download link.
) else (
    echo Build failed!
    exit /b 1
)

