# Documentação — Rate Limiting (estudo profissional)

Índice dos textos desta pasta. Leia na ordem; cada capítulo aponta para exemplos em [`../examples/`](../examples/).

| # | Documento | Exemplo relacionado |
|---|-----------|---------------------|
| 1 | [Por que rate limiting](01-por-que-rate-limiting.md) | — |
| 2 | [Conceitos e algoritmos](02-conceitos-e-algoritmos.md) | — |
| 3 | [Token bucket com x/time/rate](03-token-bucket-x-rate.md) | [`01-token-bucket-basico`](../examples/01-token-bucket-basico/main.go) |
| 4 | [Middleware e HTTP 429](04-middleware-429.md) | [`02-middleware-global`](../examples/02-middleware-global/main.go) |
| 5 | [Limite por IP e chave](05-limite-por-ip-e-chave.md) | [`03-middleware-per-ip`](../examples/03-middleware-per-ip/main.go), [`04-fixed-window-manual`](../examples/04-fixed-window-manual/main.go) |
| 6 | [Chi, Echo e produção](06-chi-echo-e-producao.md) | [`06-chi-bridge`](../examples/06-chi-bridge/main.go), [`07-echo-bridge`](../examples/07-echo-bridge/main.go) |
| 7 | [Armadilhas](07-armadilhas.md) | [`05-stdlib-server`](../examples/05-stdlib-server/main.go) |

Trilhas relacionadas: [Middleware](../../Middleware/docs/README.md), [Goroutines](../../Goroutines/docs/README.md).

Voltar ao [README da pasta Rate-Limiting](../README.md).
