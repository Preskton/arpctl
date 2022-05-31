

.PHONY: build
build:
	go build -o dist/bin/arpctl .

.PHONY: dist
dist:
	GOOS=linux GOARCH=amd64 go build -o dist/bin/arpctl-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dist/bin/arpctl-linux-arm64 .
	GOOS=linux GOARCH=arm go build -o dist/bin/arpctl-linux-arm .

.PHONY: run
run:
	@go run .

.PHONY: ship
ship: dist
	scp dist/bin/arpctl-linux-arm preston@marceline:/home/preston/arpctl
	scp config/notes.json preston@marceline:/home/preston/config
	scp config/scales.json preston@marceline:/home/preston/config
	scp config/arpconfig.yaml preston@marceline:/home/preston/config

.PHONY: run-remote
run-remote: ship
	ssh preston@marceline '/home/preston/arpctl'

.PHONY: docs
docs: build
	dist/bin/arpctl codegen docs
