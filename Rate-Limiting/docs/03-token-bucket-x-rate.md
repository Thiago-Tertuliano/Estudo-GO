# 3. Token bucket com `golang.org/x/time/rate`

## Pacote

```go
import "golang.org/x/time/rate"
```

Não faz parte da stdlib, mas é o padrão de fato no ecossistema Go (usado indiretamente em muitos projetos).

## Criar limiter

```go
// 5 eventos por segundo, burst de 10
l := rate.NewLimiter(5, 10)
```

- Primeiro argumento: `rate.Limit` (float64 = eventos/segundo). Use `rate.Every(time.Second)` para 1 req/s.
- Segundo: **burst** (capacidade do balde).

## Operações principais

| Método | Comportamento |
|--------|----------------|
| `Allow()` | Consome 1 token se houver; retorna `bool` (não bloqueia) |
| `Wait(ctx)` | Bloqueia até haver token ou `ctx` cancelar |
| `Reserve()` | Agenda token futuro (avançado) |

Para middleware HTTP, **`Allow()`** é o usual: rejeita na hora com 429.

## Exemplo sem HTTP

[`examples/01-token-bucket-basico`](../examples/01-token-bucket-basico/main.go) — loop que mostra aceites e rejeições.

```powershell
go run ./examples/01-token-bucket-basico
```

Próximo: [04-middleware-429.md](04-middleware-429.md).
