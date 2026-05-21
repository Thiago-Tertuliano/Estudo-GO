# 5. Limite por IP e chave

## Por IP

`RemoteAddr` traz `IP:porta`. Use `net.SplitHostPort` para extrair o IP.

[`internal/ratelimit/per_ip.go`](../internal/ratelimit/per_ip.go) mantém `map[string]*rate.Limiter` protegido por `sync.Mutex` — um token bucket por IP.

Exemplo: [`examples/03-middleware-per-ip`](../examples/03-middleware-per-ip/main.go).

## Por API key ou user ID

Mesma ideia: a **chave** do mapa é o valor do header `Authorization`, claim JWT (`sub`) ou tenant ID — **depois** de autenticação válida.

Ordem típica:

```text
Auth → RateLimit(por user) → Handler
```

## Janela fixa manual

[`internal/ratelimit/fixed_window.go`](../internal/ratelimit/fixed_window.go) + [`examples/04-fixed-window-manual`](../examples/04-fixed-window-manual/main.go) mostram contador por janela sem `x/time/rate`.

## Produção

| Abordagem | Quando |
|-----------|--------|
| Map em memória | Protótipo, single instance, IPs limitados |
| Redis + TTL | Cluster, muitos IPs, janelas compartilhadas |
| API Gateway | Rate limit na borda, app sem estado |

Próximo: [06-chi-echo-e-producao.md](06-chi-echo-e-producao.md).
