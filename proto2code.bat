@echo off
REM ===== proto2code.bat =====
REM Generates Go gRPC code into protos/openinventorymanager

SET PROTO_PATH=protos
SET PROTO_FILE=openinventorymanager.proto
SET GO_OUT=.

echo Generating Go code from %PROTO_PATH%\%PROTO_FILE% ...

protoc --proto_path=%PROTO_PATH% --go_out=%GO_OUT% --go-grpc_out=%GO_OUT% %PROTO_FILE%

IF %ERRORLEVEL% NEQ 0 (
    echo Failed to generate proto code!
    exit /b %ERRORLEVEL%
)

echo Proto generation successful!
pause