# AGENTS

Este documento padroniza a forma de usar agentes de IA na evolução do projeto. Seu propósito é garantir que toda automação siga a arquitetura e os princípios definidos.

## Exemplo de Prompt
```
Crie um novo módulo chamado `Invoice` seguindo os mesmos padrões do módulo `User`:
- `entity`
- `repository` (interface)
- `service` (interface)
- `usecase` (implementação)
- `handler` com endpoints REST
- Registro de rotas e injeção no container
```

## Checklist de Validação SOLID
- A responsabilidade de cada struct está clara e única?
- Alguma interface está sendo implementada forçadamente?
- A lógica de domínio conhece detalhes técnicos da infraestrutura?
- Algum handler ou usecase realiza mais de uma função?
- As dependências são injetadas por abstrações?

Regras
- O tom deve ser direto e técnico
- Não incluir código, apenas estrutura, instruções e convenções

⸻

Padrões aplicados
- SRP: cada documento tem uma função clara
- Transparência e padronização para evoluções futuras

Finalidade
Documentar com clareza a estrutura, funcionamento e diretrizes do projeto, tornando possível sua adoção por múltiplos desenvolvedores e agentes de automação com mínima curva de aprendizado.
