generate:
	@protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

build: generate
	@go build -ldflags "-s -w" -o bin/prototut ./server/...
	@upx --lzma bin/prototut

run: generate build
	@bin/prototut -port=8080

dev: generate
	@go run ./server/... --port=8080

clean:
	@rm -rf bin
