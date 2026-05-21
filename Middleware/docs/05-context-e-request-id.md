# 5. Context e request ID

## `r.Context()`

Cada request traz um `context.Context` (cancelável quando o cliente desconecta). Middleware pode enriquecer:

```go
ctx := context.WithValue(r.Context(), key, value)
next.ServeHTTP(w, r.WithContext(ctx))
```

## Chaves tipadas

Evite `context.WithValue(ctx, "requestID", id)` com string solta — colisão com outras libs. Use tipo privado:

```go
type ctxKey int
const requestIDKey ctxKey = 1
```

Leitura no handler:

```go
id, ok := httpmw.GetRequestID(r.Context())
```

## Header `X-Request-ID`

Middleware profissional:

1. Se o cliente enviou `X-Request-ID`, reutiliza.
2. Senão, gera ID (UUID ou random hex).
3. Define no response header e no context.

## Exemplo

[`06-context-value`](../examples/06-context-value/main.go) — pacote [`internal/httpmw/request_id.go`](../internal/httpmw/request_id.go).

## Não abusar de `WithValue`

Context é para **dados de request-scoped** (ID, claims de auth já validados), não para passar dependências grandes (DB, config) — use injeção no handler ou struct de app.

Próximo: [06-chi-echo-vs-stdlib.md](06-chi-echo-vs-stdlib.md).
