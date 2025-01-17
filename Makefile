semversort-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o semversort-linux-amd64

semversort-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o semversort-darwin-amd64

.PHONY: build
build:
	go build ./...
