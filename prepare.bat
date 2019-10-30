protoc -I ./proto --go_out=plugins=grpc:./proto transfer-service.proto
setlocal
    set GOOS=linux
    set GOARCH=amd64
    set CGO_ENABLED=0
    go build -o tfc-transfer-validator -a -installsuffix cgo ./main/main.go
endlocal
docker build -t tfc/tfc-transfer-validator .