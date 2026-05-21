# 4. ResponseWriter wrapper

## Problema

`http.ResponseWriter` não expõe o **status** já enviado de forma confiável antes do fim da resposta. Middlewares de log precisam interceptar `WriteHeader`.

## Solução: embedding

```go
type StatusRecorder struct {
    http.ResponseWriter
    Status int
}

func (r *StatusRecorder) WriteHeader(code int) {
    r.Status = code
    r.ResponseWriter.WriteHeader(code)
}
```

Passe `rec` para `next.ServeHTTP(rec, r)` no middleware de logging.

## Implementação no repo

[`internal/httpmw/status_writer.go`](../internal/httpmw/status_writer.go) e uso em [`logging.go`](../internal/httpmw/logging.go).

## Exemplo dedicado

[`04-status-recorder`](../examples/04-status-recorder/main.go)

## Cuidado

Se o handler nunca chamar `WriteHeader`, o status efetivo é **200** na primeira `Write` — o recorder deve tratar isso (como em `Write` do pacote `httpmw`).

Próximo: [05-context-e-request-id.md](05-context-e-request-id.md).
