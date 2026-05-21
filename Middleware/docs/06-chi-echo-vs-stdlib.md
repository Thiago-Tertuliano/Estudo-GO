# 6. Chi, Echo e stdlib

## Mesma ideia, assinaturas diferentes

| Camada | Registro | Tipo do middleware |
|--------|----------|-------------------|
| **net/http** | `handler = mw(final)` | `func(http.Handler) http.Handler` |
| **Chi** | `r.Use(mw)` | `func(http.Handler) http.Handler` (compatível com stdlib) |
| **Echo** | `e.Use(mw)` | `func(echo.HandlerFunc) echo.HandlerFunc` |

Chi aceita middlewares escritos para stdlib **diretamente** — por isso `r.Use(middleware.Logger)` do pacote chi funciona.

Echo usa assinatura própria; para reutilizar lógica stdlib, use **adaptador** (ver exemplo 10).

## Seus projetos

**Chi** — [Curso-API-GO/main.go](../../Estudos-Realizados/Curso-API-GO/main.go):

```go
r.Use(middleware.Logger)
r.Use(middleware.Recoverer)
r.Use(middleware.RequestID)
```

**Echo** — [fitness-api/main.go](../../Estudos-Realizados/API-Kelche/fitness-api/main.go):

```go
e.Use(middleware.Logger())
e.Use(middleware.CORSWithConfig(...))
```

Depois desta trilha, você pode abrir o código-fonte desses middlewares no GitHub e mapear: wrapper de `ResponseWriter`, `recover`, etc.

## Exemplos desta trilha

- [`09-chi-bridge`](../examples/09-chi-bridge/main.go) — `httpmw.Logging` + `chi.Use`
- [`10-echo-bridge`](../examples/10-echo-bridge/main.go) — adaptador stdlib → Echo

Próximo: [07-armadilhas.md](07-armadilhas.md).
