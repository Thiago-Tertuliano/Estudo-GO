# 2. Assinatura do middleware

## Definição

Middleware HTTP em Go (stdlib) é uma função que **recebe** o próximo handler e **retorna** um handler envolvido:

```go
func MeuMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // antes do handler
        next.ServeHTTP(w, r)
        // depois do handler
    })
}
```

Alias comum de tipo:

```go
type Middleware func(http.Handler) http.Handler
```

## Não confundir com o handler final

| Papel | Assinatura típica |
|-------|-------------------|
| **Handler final** | `func(w, r)` — responde ao cliente |
| **Middleware** | `func(next http.Handler) http.Handler` — delega com `next.ServeHTTP` |

Esquecer `next.ServeHTTP(w, r)` faz a requisição **nunca** chegar ao handler interno.

## Onde registrar

```go
handler := MeuMiddleware(finalHandler)
mux.Handle("/", handler)
```

Ou empilhar vários (ver doc 3).

## Exemplo

[`02-wrap-handler`](../examples/02-wrap-handler/main.go)

Próximo: [03-cadeia-e-ordem.md](03-cadeia-e-ordem.md).
