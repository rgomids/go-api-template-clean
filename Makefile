.PHONY: help run dev lint fmt test build docker-dev docker-prod setup

help:
	@echo "Comandos disponíveis:"
	@echo "  make run          - Executa a API localmente (modo padrão)"
	@echo "  make dev          - Executa a API com variável ENV=dev"
	@echo "  make setup        - Prepara o ambiente de desenvolvimento"
	@echo "  make test         - Executa os testes unitários"
	@echo "  make lint         - Roda o linter (go vet + staticcheck)"
	@echo "  make fmt          - Formata o código com go fmt"
	@echo "  make build        - Compila o binário para o host local"
	@echo "  make docker-dev   - Builda a imagem Docker de dev"
	@echo "  make docker-prod  - Builda a imagem Docker de produção"

run:
	go run ./cmd/main.go

dev:
	ENV=dev go run ./cmd/main.go

setup:
	@cp -n .env.example .env 2>/dev/null || true
	go mod download
	go install honnef.co/go/tools/cmd/staticcheck@latest

test:
	go test ./... -v -cover

lint:
	go vet ./...
	staticcheck -go 1.24 ./...

fmt:
	go fmt ./...

build:
	go build -o bin/api ./cmd/main.go

docker-dev:
	docker build -f infra/docker/Dockerfile.dev -t go-api-template:dev .

docker-prod:
	docker build -f infra/docker/Dockerfile.prod -t go-api-template:prod .
