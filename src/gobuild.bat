@echo off
@echo on
go clean -v ./...
go build -o . -v ./...
go run ./cmd/restserveur/main.go
@echo off
if %ERRORLEVEL% GEQ 1 echo !!!!! ERROR !!!!!