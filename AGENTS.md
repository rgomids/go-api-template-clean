# AGENTS

## Propósito

Orientar agentes de IA a evoluírem o projeto sem comprometer a arquitetura limpa e os padrões estabelecidos.

## Uso da CLI de scaffolding

Sempre utilize a ferramenta `go-api-cli` para criar novas entidades. Não edite arquivos manualmente.

Exemplo de comando:

```bash
make scaffold entity=Invoice fields="number:string total:float status:enum[pending,paid,failed] user:belongsTo"
```

O scaffold cria código em todas as camadas, testes com cobertura total e migrations.
Para compilar a CLI execute:

```bash
make build-cli
```

## Checklist SOLID

- Cada struct possui responsabilidade única?
- Interfaces são pequenas e específicas?
- O domínio não importa pacotes externos ou de infra?
- Os testes cobrem todos os fluxos esperados?
- As migrations refletem os tipos e restrições definidos?

### Padrões reforçados
- SRP em cada arquivo gerado
- DIP garantido por interfaces e mocks
- Test First através dos templates
- DRY via geração automática

Sempre valide as alterações rodando:

```bash
go test ./...
```
