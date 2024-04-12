
build:
	go build -v -o auth-service.out ./cmd/service/main.go

build_stripped:
	go build -ldflags="-w -s" -o auth-service.out ./cmd/service/main.go

test:
	go test ./...

makemigration:
	migrate create -ext sql -dir migrations $(name)

.DEFAULT_GOAL := build
