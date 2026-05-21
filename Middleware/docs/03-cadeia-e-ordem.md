# 3. Cadeia e ordem

## Modelo "cebola"

Request entra pelo middleware **mais externo**; response sai de volta pelas camadas que têm código **após** `next.ServeHTTP`.

Registro típico (primeiro `Use` = mais externo):

```text
Recover → RequestID → Logging → Handler
```

Implementação com loop (ver `internal/httpmw/chain.go`):

```go
func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
    for i := len(mws) - 1; i >= 0; i-- {
        h = mws[i](h)
    }
    return h
}
```

## Ordem importa

- **Recover** deve ficar **por fora** para capturar panic de middlewares internos e do handler.
- **Logging** costuma ficar fora do handler mas dentro do recover, para logar status final.
- **Auth** antes do handler de negócio; depois de recover se quiser logar falhas de auth.

## Global vs rota

- **Global:** `mux.Handle("/", Chain(handler, mw1, mw2))` ou wrapper no `ListenAndServe`.
- **Por grupo:** Chi `r.Route("/api", func(r chi.Router) { r.Use(auth); ... })`.

## Exemplos

- [`05-chain`](../examples/05-chain/main.go)
- [`08-stdlib-server`](../examples/08-stdlib-server/main.go)

Próximo: [04-responsewriter-wrapper.md](04-responsewriter-wrapper.md).
