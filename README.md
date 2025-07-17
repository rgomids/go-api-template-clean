# go-api-template-clean

Este repositório apresenta uma estrutura base para APIs em Go.
Consulte o diretório [`docs`](docs/README.md) para detalhes de arquitetura e instruções de uso.

## Configuração rápida

1. Copie o arquivo `.env.example` para `.env` e ajuste os valores conforme o ambiente:
   ```bash
   cp .env.example .env
   ```
   Variáveis disponíveis:
   - `APP_ENV` (dev|prod)
   - `PORT` (padrão 8080)
   - `DATABASE_URL` (obrigatória)
   - `REDIS_URL`
   - `SMTP_HOST`
   - `SMTP_PORT` (padrão 587)
   - `SMTP_USER`
   - `SMTP_PASSWORD`
2. Instale as dependências e execute a aplicação:
   ```bash
   go mod download
   go run ./cmd
   ```
