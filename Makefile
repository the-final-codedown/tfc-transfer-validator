build:
	protoc -I. --gofast_out=plugins=grpc:. \
		proto/tfc-transfer-validator.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o tfc-transfer-validator -a -installsuffix cgo ./main/main.go
	docker build -t tfc/tfc-transfer-validator .

local:
	protoc -I. --gofast_out=plugins=grpc:. \
		proto/tfc-transfer-validator.proto
	go build -o tfc-transfer-validator -a -installsuffix cgo ./main/main.go
	docker build -t tfc/tfc-transfer-validator .

run:
	docker run -p 50052:50052 -e SERVER_PORT=:50052 tfc/tfc-transfer-validator