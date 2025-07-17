# AGENTS

Este arquivo orienta o uso de IA na expansão do projeto.

## Exemplo de prompt para criação de novos módulos

```
Crie um novo módulo de domínio seguindo a estrutura do projeto. Gere entidade, repositório, serviço, usecase e handler correspondentes, aplicando princípios SOLID e mantendo a independência de infraestrutura.
```

## Checklist de validação SOLID

- [ ] **Single Responsibility**: cada módulo possui apenas uma responsabilidade?
- [ ] **Open/Closed**: novas funcionalidades exigem modificação mínima em código existente?
- [ ] **Liskov Substitution**: interfaces são respeitadas por implementações?
- [ ] **Interface Segregation**: dependências expõem apenas o necessário?
- [ ] **Dependency Inversion**: detalhes dependem de abstrações e não o contrário?
