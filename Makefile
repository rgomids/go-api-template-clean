.PHONY: help run dev lint fmt test coverage build docker-dev docker-prod setup build-cli scaffold up down logs ps

CLI_BIN := ./bin/go-api-cli

help:
	@echo "Comandos disponíveis:"
	@echo "  make run          - Executa a API localmente (modo padrão)"
	@echo "  make dev          - Executa a API com variável ENV=dev"
	@echo "  make setup        - Prepara o ambiente de desenvolvimento"
	@echo "  make test         - Executa os testes unitários"
	@echo "  make coverage     - Executa testes com relatório de cobertura"
	@echo "  make lint         - Roda o linter (go vet + staticcheck)"
	@echo "  make fmt          - Formata o código com go fmt"
	@echo "  make build        - Compila o binário para o host local"
	@echo "  make docker-dev   - Builda a imagem Docker de dev"
	@echo "  make docker-prod  - Builda a imagem Docker de produção"
	@echo "  make build-cli    - Compila a CLI de scaffolding"
	@echo "  make scaffold     - Gera código de entidade via CLI"
	@echo "  make up           - Sobe dependências via docker-compose"
	@echo "  make down         - Encerra containers do docker-compose"
	@echo "  make logs         - Visualiza logs dos serviços"
	@echo "  make ps           - Lista status dos serviços"

run:
	go run ./cmd/main.go

dev:
	ENV=dev go run ./cmd/main.go

setup: build build-cli
	@cp -n .env.example .env 2>/dev/null || true
	go mod download
	go install honnef.co/go/tools/cmd/staticcheck@latest

test:
	go test ./... -v -cover

coverage:
	mkdir -p coverage
	go test ./... -coverprofile=coverage/coverage.out
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html

lint:
	go vet ./...
	staticcheck -go 1.24 ./...

fmt:
	go fmt ./...

build:
	go build -o bin/api ./cmd/main.go

build-cli:
	go build -o $(CLI_BIN) ./cli/cmd/main.go

scaffold:
	$(CLI_BIN) scaffold $(entity) $(fields)
	$(MAKE) fmt

docker-dev:
	docker build -f infra/docker/Dockerfile.dev -t go-api-template:dev .

docker-prod:
	docker build -f infra/docker/Dockerfile.prod -t go-api-template:prod .

up:
	docker compose -f docker/dev/docker-compose.yml up -d

down:
	docker compose -f docker/dev/docker-compose.yml down

logs:
	docker compose -f docker/dev/docker-compose.yml logs -f

ps:
	docker compose -f docker/dev/docker-compose.yml ps
