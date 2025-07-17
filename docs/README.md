# Go API Template Clean

## Visão Geral
Este repositório fornece um ponto de partida para a criação de APIs em Go voltadas para escalabilidade e testes. Foi desenhado para desenvolvedores que desejam aplicar Clean Architecture e princípios SOLID desde os primeiros commits.

## Arquitetura
A estrutura segue os conceitos da Clean Architecture onde regras de negócio ficam isoladas das camadas externas. Cada pacote expõe apenas aquilo que a camada superior necessita, favorecendo o uso de dependências injetadas por interfaces.

- **domain**: contém entidades e regras de negócio.
- **infra**: implementa detalhes técnicos (banco, cache, email).
- **handler**: camada de entrada, atualmente via HTTP.
- **app**: compõe e injeta dependências entre as camadas.

Essas separações reforçam os princípios SOLID, mantendo baixo acoplamento e alta coesão.

## Padrões Utilizados
- Factory
- Adapter
- Strategy
- Middleware
- Functional Options (planejado para usos futuros)
- Interface Segregation
- Dependency Inversion

## Estrutura do Projeto
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
Cada pasta possui responsabilidade única, facilitando testes unitários e extensões controladas.

## Como Usar
1. Clone o projeto e instale as dependências:
   ```bash
git clone <repo-url>
cd go-api-template-clean
go mod download
   ```
2. Configure as variáveis de ambiente conforme `internal/config`:
   - `APP_ENV` (dev|prod)
   - `PORT` (padrão 8080)
   - `DATABASE_URL` (obrigatória)
   - `REDIS_URL`
   - `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASSWORD`
3. Execute a aplicação:
   ```bash
go run ./cmd
   ```
   Quando um `Makefile` estiver disponível utilize `make build`, `make test` e `make run` para padronizar as etapas.

## Como Estender
- Crie novas entidades em `internal/domain/entity` e suas interfaces em `repository` ou `service`.
- Implemente casos de uso em `usecase` mantendo a orquestração da lógica de negócio.
- Adicione handlers e rotas em `internal/handler/http` e registre-os em `routes`.
- Para cada nova dependência externa, forneça uma implementação em `internal/infra` e injete-a via `app`.
- Utilize os padrões e a checklist de `docs/AGENTS.md` para garantir aderência aos princípios SOLID.
