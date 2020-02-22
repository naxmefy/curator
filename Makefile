.PHONY: install

install:
	go install ./...

test:
	go test -cover ./...

secrets:
	go run cmd/secrets/main.go ./...