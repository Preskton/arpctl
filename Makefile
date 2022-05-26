

.PHONY: build
build:
	go build ./...

.PHONY: dist
dist:
	GOOS=linux GOARCH=amd64 go build -o dist/bin/arpctl-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o dist/bin/arpctl-linux-arm64
	GOOS=linux GOARCH=arm go build -o dist/bin/arpctl-linux-arm
