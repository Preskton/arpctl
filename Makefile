

.PHONY: build
build:
	go build -o dist/bin/arpctl

.PHONY: dist
dist:
	GOOS=linux GOARCH=amd64 go build -o dist/bin/arpctl-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o dist/bin/arpctl-linux-arm64
	GOOS=linux GOARCH=arm go build -o dist/bin/arpctl-linux-arm

.PHONY: run
run:
	@go run .

ship: dist
	scp dist/bin/arpctl-linux-arm preston@marceline:/home/preston/arpctl

run-remote:
	ssh preston@marceline '/home/preston/arpctl'
