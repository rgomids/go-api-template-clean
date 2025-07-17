# go-api-template-clean

Este repositório apresenta uma estrutura base para APIs em Go.
Consulte o diretório [`docs`](docs/README.md) para detalhes de arquitetura e instruções de uso.

## Requisitos

- Go 1.20 ou superior
- Make
- Ferramentas de lint como [staticcheck](https://staticcheck.io) (instalado via `make setup`)

## Conexões externas

O projeto já inclui integrações para:

- Banco de dados PostgreSQL
- Cache Redis
- Envio de emails via SMTP

## Configuração rápida

1. Execute o comando abaixo para preparar o ambiente de desenvolvimento:
   ```bash
   make setup
   ```
   Esse passo copia o `.env.example` para `.env` (caso ainda não exista), instala as dependências Go e a ferramenta `staticcheck`.
2. Ajuste as variáveis no arquivo `.env`:
   - `APP_ENV` (dev|prod)
   - `PORT` (padrão 8080)
   - `DATABASE_URL` (obrigatória)
   - `REDIS_URL`
   - `SMTP_HOST`
   - `SMTP_PORT` (padrão 587)
   - `SMTP_USER`
   - `SMTP_PASSWORD`
3. Execute a aplicação:
   ```bash
   make run
   ```

## Testes e cobertura

Para executar os testes com relatório de cobertura, utilize:

```bash
make coverage
```
O comando gera os arquivos `coverage/coverage.out` e `coverage/coverage.html`.

## Rotas disponíveis

- `GET /health` retorna o status e a versão da API.
- `POST /users` cria um usuário.
- `DELETE /users/{id}` remove um usuário.

Importe a coleção `docs/postman_collection.json` no Postman para testar todos os endpoints rapidamente.
