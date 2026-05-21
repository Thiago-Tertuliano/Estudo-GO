# 2. Conceitos e algoritmos

## Vocabulário

| Termo | Significado |
|-------|-------------|
| **Taxa (rate)** | Eventos permitidos por unidade de tempo (ex.: 10/s) |
| **Burst** | Rajada máxima aceita de uma vez (token bucket) |
| **Janela** | Intervalo em que a contagem vale (ex.: 1 minuto) |
| **Chave** | Identificador do bucket: IP, API key, user ID |

## Algoritmos comuns

### Token bucket (usado nesta trilha via `x/time/rate`)

- Balde com tokens que enchem a taxa `r` por segundo.
- Cada requisição consome 1 token se `Allow()`; `burst` define quantos tokens cabem no balde.
- Permite **picos curtos** (burst) mantendo média no longo prazo.

### Fixed window (janela fixa)

- Contador zera a cada janela (ex.: minuto calendário ou minuto desde primeira req).
- Simples de implementar com `map` + `mutex` — exemplo [`04-fixed-window-manual`](../examples/04-fixed-window-manual/main.go).
- **Problema:** na virada da janela o cliente pode enviar `2× max` em poucos segundos.

### Sliding window / leaky bucket

- Mais justos; implementação mais trabalhosa.
- Em produção distribuída costuma-se Redis com scripts Lua ou serviços gerenciados.

## Escolha rápida

| Cenário | Sugestão |
|---------|----------|
| API Go monólito, limite global | `rate.NewLimiter` |
| Limite por IP em estudo local | `PerIP` neste repo (com ressalvas) |
| Cluster com N réplicas | Store compartilhado (Redis), não map em memória |
| Login / endpoint sensível | Limite **por rota** + chave (user/IP) mais restrito |

Próximo: [03-token-bucket-x-rate.md](03-token-bucket-x-rate.md).
