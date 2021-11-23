@echo off
setlocal
pushd "%~dp0"

set _ROOT_=%CD%\..
set _RES_=%CD%\res.txt

cd %_ROOT_%
type "%~f1"|go run main.go "%~2" > "%_RES_%"

popd
endlocal
exit /b
