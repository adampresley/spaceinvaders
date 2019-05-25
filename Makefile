.DEFAULT_GOAL := run

generate:
	go generate

build-mac: generate
	GOOS=darwin GOARCH=amd64 go build

build-windows: generate
	GOOS=windows GOARCH=amd64 go build

build-linux: generate
	GOOS=linux GOARCH=amd64 go build

run:
	go run *.go