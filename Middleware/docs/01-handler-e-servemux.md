# 1. Handler e ServeMux

## Interface `http.Handler`

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Qualquer tipo com método `ServeHTTP` pode atender HTTP. O mais comum é `http.HandlerFunc` — função com essa assinatura:

```go
http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ok")
})
```

## Servidor mínimo

```go
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil) // DefaultServeMux
```

Ou mux explícito:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /{$}", home) // Go 1.22+ padrões de rota
http.ListenAndServe(":8080", mux)
```

## Request e ResponseWriter

- `*http.Request` — método, URL, headers, body, **`Context()`** (cancelamento por cliente).
- `http.ResponseWriter` — escreve status, headers e body da resposta.

Cada request é tratada em **uma goroutine** pelo `net/http` (ver [Goroutines](../../Goroutines/docs/02-context-e-cancelamento.md)).

## Exemplos

- [`01-hello-handler`](../examples/01-hello-handler/main.go)
- [`08-stdlib-server`](../examples/08-stdlib-server/main.go)

Próximo: [02-assinatura-middleware.md](02-assinatura-middleware.md).
