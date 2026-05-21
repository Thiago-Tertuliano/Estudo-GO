# 4. Middleware e HTTP 429

## Padrão

Mesma assinatura dos outros middlewares:

```go
func Middleware(l *rate.Limiter) func(http.Handler) http.Handler
```

Implementação em [`internal/ratelimit/limiter.go`](../internal/ratelimit/limiter.go):

1. `if !l.Allow()` → responde **429**, **não** chama `next.ServeHTTP`.
2. Caso contrário, delega ao handler.

## Headers

| Header | Uso |
|--------|-----|
| `Retry-After` | Sugestão em segundos (ex.: `"1"`) |
| `X-RateLimit-Limit` | Taxa configurada (informativo) |

## Ordem na chain

```text
Recover → Logging → RateLimit → Handler
```

Recover por fora para capturar panic; rate limit **antes** do trabalho pesado do handler.

## Exemplo

[`examples/02-middleware-global`](../examples/02-middleware-global/main.go) — limite global baixo (fácil testar com curl em loop).

[`examples/05-stdlib-server`](../examples/05-stdlib-server/main.go) — servidor completo.

Próximo: [05-limite-por-ip-e-chave.md](05-limite-por-ip-e-chave.md).
