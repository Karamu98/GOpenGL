go build
xcopy "%~dp0res" "%~dp0bin\res" /s /y /i /c
move "%~dp0Project.exe" "%~dp0bin"
