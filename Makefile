all: x86_64 armv7

x86_64: server.go
	mkdir -p build
	env GOOS=linux GOARCH=amd64 go build  -o build/server.x86_64 server.go

armv7: server.go
	mkdir -p build
	env GOOS=linux GOARCH=arm GOARM=7 go build -o build/server.armv7 server.go
