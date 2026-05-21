# 6. Chi, Echo e produção

## Stdlib e Chi — mesma assinatura

Chi `r.Use` aceita `func(http.Handler) http.Handler`:

```go
r := chi.NewRouter()
limiter := rate.NewLimiter(2, 5)
r.Use(ratelimit.Middleware(limiter))
```

Exemplo: [`examples/06-chi-bridge`](../examples/06-chi-bridge/main.go).

## Echo — adaptador

Echo usa `func(echo.HandlerFunc) echo.HandlerFunc`. Reutilize o padrão de [`Middleware/examples/10-echo-bridge`](../../Middleware/examples/10-echo-bridge/main.go):

```go
func wrapStd(mw func(http.Handler) http.Handler) echo.MiddlewareFunc { ... }
e.Use(wrapStd(ratelimit.Middleware(limiter)))
```

Exemplo: [`examples/07-echo-bridge`](../examples/07-echo-bridge/main.go).

## Seus projetos hoje

| Projeto | Rate limit hoje |
|---------|-----------------|
| [Curso-API-GO](../../Estudos-Realizados/Curso-API-GO/main.go) | Não (só Logger, Recoverer, RequestID, Timeout) |
| [API-Kelche](../../Estudos-Realizados/API-Kelche/fitness-api/README.md) | Listado como melhoria futura |

Após esta trilha, um passo natural é adicionar `r.Use(ratelimit.Middleware(...))` no Chi ou equivalente no Echo em ambiente de homologação.

## Gateway vs app

- **Gateway** (Kong, NGINX limit_req): protege antes de chegar no Go.
- **App**: quotas por usuário/plano, regras finas por rota.

Os dois podem coexistir.

Próximo: [07-armadilhas.md](07-armadilhas.md).
