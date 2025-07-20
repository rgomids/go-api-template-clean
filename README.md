# go-api-template-clean

Este repositório oferece uma base para construir APIs em Go seguindo Clean Architecture e princípios SOLID.

## Requisitos

- Go 1.20 ou superior
- Make
- Ferramentas de lint como [staticcheck](https://staticcheck.io) (instaladas via `make setup`)

## Visão Geral da Arquitetura

A estrutura divide responsabilidades em camadas bem definidas:

- **domain**: entidades e regras de negócio
- **infra**: detalhes de infraestrutura (banco, cache, email)
- **handler**: entrada HTTP
- **app**: composição e injeção de dependências

Essas separações reforçam baixo acoplamento e alta coesão.

### Padrões Utilizados
- Factory
- Adapter
- Strategy
- Middleware
- Interface Segregation
- Dependency Inversion

### Estrutura de Diretórios
```text
.
├── cmd/                  ponto de entrada da aplicação
├── docs/                 documentação e guias
├── internal/
│   ├── app/              injeção de dependências
│   ├── config/           carregamento de variáveis de ambiente
│   ├── domain/           regras de negócio
│   │   ├── entity/       modelos de domínio
│   │   ├── repository/   contratos de persistência
│   │   ├── service/      interfaces de serviços
│   │   └── usecase/      orquestração de regras
│   ├── handler/          camadas de entrada (HTTP)
│   │   └── http/         handlers, rotas e middlewares
│   └── infra/            implementações concretas (db, cache, email)
└── pkg/                  utilidades e contratos externos
```

## Conexões externas

O projeto já contempla integrações para PostgreSQL, Redis e SMTP.

## Configuração rápida

1. Execute:
   ```bash
   make setup
   ```
   Esse passo copia `.env.example` para `.env` (caso não exista), instala dependências Go e `staticcheck`.
2. Ajuste as variáveis no `.env` conforme `internal/config`.
3. Rode a aplicação:
   ```bash
   make run
   ```

### 🔧 Subir ambiente de desenvolvimento

```bash
   make up
```

Banco de dados e Redis serão inicializados com:
- PostgreSQL: localhost:5432
- Redis: localhost:6379

Os valores podem ser ajustados copiando `.env.example` para `.env` e editando as
variáveis `DB_*` e `REDIS_*`. O `docker-compose` lerá essas variáveis para
configurar os serviços.

Para encerrar:

```bash
   make down
```

## Testes e cobertura

```bash
make coverage
```
Gera `coverage/coverage.out` e `coverage/coverage.html`.

## 🔧 Scaffold de novas entidades

Este projeto inclui a CLI `go-api-cli` para gerar scaffolds completos.

Exemplo:
```bash
make scaffold entity=Book fields="title:string author:string genre:enum[fiction,non-fiction] pages:int"
```
O comando cria arquivos em todas as camadas, testes automatizados e migrations. Para compilar a CLI:
```bash
make build-cli
```

## Rotas disponíveis

- `GET /health` retorna o status da API
- `POST /users` cria um usuário
- `DELETE /users/{id}` remove um usuário

Importe `docs/postman_collection.json` e `docs/postman_environment.json` no Postman para testar.

## Como estender

- Crie entidades em `internal/domain/entity` e interfaces em `repository` ou `service`.
- Implemente casos de uso em `usecase`.
- Adicione handlers e rotas em `internal/handler/http` e registre-as em `routes`.
- Para cada dependência externa, forneça uma implementação em `internal/infra` e injete via `app`.
- Consulte [AGENTS.md](AGENTS.md) para garantir aderência aos princípios SOLID.

### Referências
- [Wikipedia - Software design pattern](https://en.wikipedia.org/wiki/Software_design_pattern)
- [Refactoring Guru - Design Patterns in Go](https://refactoring.guru/design-patterns/go)
