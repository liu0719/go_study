@REM 先设为linux
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o main main.go

@REM 再改回windows
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -o main.exe main.go
@REM linux 下运行命令，先把main 改为可执行,再运行
@REM chmod +x main
@REM ./main