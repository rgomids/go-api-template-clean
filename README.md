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

## Rotas disponíveis

- `GET /health` retorna o status e a versão da API.
- `POST /users` cria um usuário.
- `DELETE /users/{id}` remove um usuário.

Importe a coleção `docs/postman_collection.json` no Postman para testar todos os endpoints rapidamente.

## Evoluindo a API

Siga os passos abaixo para adicionar novas funcionalidades. O exemplo a seguir mostra como criar uma rota para cadastro de produtos.

1. Crie a estrutura do produto em `internal/domain/entity/product.go`:
   ```go
   package entity

   type Product struct {
       ID   string
       Name string
   }
   ```
2. Defina `ProductRepository` em `internal/domain/repository` e a interface `ProductService` em `internal/domain/service`.
3. Implemente `ProductUseCase` em `internal/domain/usecase` seguindo as interfaces criadas.
4. Crie `ProductHandler` em `internal/handler/http` para expor os métodos via HTTP.
5. Registre as rotas em `internal/handler/http/routes/routes.go`:
   ```go
   func RegisterRoutes(router *chi.Mux, userHandler *http.UserHandler, productHandler *http.ProductHandler) {
       router.Route("/users", func(r chi.Router) {
           r.Post("/", userHandler.Register)
           r.Delete("/{id}", userHandler.Delete)
       })

       router.Route("/products", func(r chi.Router) {
           r.Post("/", productHandler.Create)
           r.Get("/{id}", productHandler.FindByID)
       })
   }
   ```
6. Atualize `BuildContainer` em `internal/app/container.go` para injetar o novo handler.
7. Execute `go test ./...` para garantir que tudo continua funcionando.

### Referências sobre Patterns
- [Wikipedia - Software design pattern](https://en.wikipedia.org/wiki/Software_design_pattern)
- [Refactoring Guru - Design Patterns in Go](https://refactoring.guru/design-patterns/go)
