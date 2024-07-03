BINARY_NAME=bin/api-server
SOURCE_PATH=./cmd/api-server

build-windows:
	@GOOS=windows GOARCH=arm64 go build -o ${BINARY_NAME}.exe ${SOURCE_PATH}

build-linux:
	@GOOS=linux GOARCH=arm64 go build -o ${BINARY_NAME} ${SOURCE_PATH}	

run:
	@go run cmd/api-server/main.go

test:
	@go test -v ./..	