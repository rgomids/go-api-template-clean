# Go API Template Clean

## Objetivo do template

Este projeto provê uma base para construção de APIs em Go utilizando Clean Architecture e princípios SOLID. A estrutura visa facilitar a evolução modular e independente de infraestrutura, além de permitir testes unitários isolados.

## Design Patterns Principais

- Dependency Injection
- Repository
- Service
- Use Case (Interactor)
- Middleware como Função

## Arquitetura

O código é organizado no diretório `internal`, separando domínio, camadas de aplicação e infraestrutura. Use cases interagem com repositórios através de interfaces, garantindo que mudanças em bancos ou serviços externos não afetem o domínio. Handlers expõem as funcionalidades via HTTP e utilizam middlewares compostáveis.

```
cmd/            # ponto de entrada da aplicação
internal/
  config/       # carregamento de configurações
  domain/       # regras de negócio e entidades
  handler/      # handlers HTTP
  usecase/      # casos de uso
  repository/   # interfaces de repositório
  service/      # serviços de domínio
  infra/        # implementações de infraestrutura (db, email, cache)
  routes/       # roteamento da aplicação
  middleware/   # middlewares HTTP
  app/          # container de dependências
pkg/            # utilidades e contratos externos
```

## Como iniciar o projeto

1. Crie um módulo Go:

```bash
go mod init github.com/rgomids/go-api-template-clean
```

2. Construa o container de dependências e inicialize o servidor:

```bash
go run ./cmd
```

## Como estender

- Adicione novos casos de uso em `internal/usecase`
- Implemente repositórios concretos em `internal/infra` e referencie interfaces em `internal/repository`
- Utilize middlewares em `internal/middleware` compondo as rotas em `internal/routes`

## Criando novos módulos

Siga o padrão existente para adicionar novas funcionalidades. Utilize o arquivo `docs/AGENTS.md` como referência para prompts de criação automática e checklist de aderência aos princípios SOLID.
