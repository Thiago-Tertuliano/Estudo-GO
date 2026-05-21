# Documentação — Middleware HTTP (estudo profissional)

Índice dos textos desta pasta. Leia na ordem; cada capítulo aponta para exemplos em [`../examples/`](../examples/).

| # | Documento | Exemplo relacionado |
|---|-----------|---------------------|
| 1 | [Handler e ServeMux](01-handler-e-servemux.md) | [`01-hello-handler`](../examples/01-hello-handler/main.go), [`08-stdlib-server`](../examples/08-stdlib-server/main.go) |
| 2 | [Assinatura do middleware](02-assinatura-middleware.md) | [`02-wrap-handler`](../examples/02-wrap-handler/main.go) |
| 3 | [Cadeia e ordem](03-cadeia-e-ordem.md) | [`05-chain`](../examples/05-chain/main.go), [`08-stdlib-server`](../examples/08-stdlib-server/main.go) |
| 4 | [ResponseWriter wrapper](04-responsewriter-wrapper.md) | [`04-status-recorder`](../examples/04-status-recorder/main.go) |
| 5 | [Context e request ID](05-context-e-request-id.md) | [`06-context-value`](../examples/06-context-value/main.go) |
| 6 | [Chi, Echo e stdlib](06-chi-echo-vs-stdlib.md) | [`09-chi-bridge`](../examples/09-chi-bridge/main.go), [`10-echo-bridge`](../examples/10-echo-bridge/main.go) |
| 7 | [Armadilhas](07-armadilhas.md) | [`07-recover-panic`](../examples/07-recover-panic/main.go) |

Trilhas relacionadas: [Goroutines](../../Goroutines/docs/README.md) (context por request), [Channels](../../Channels/docs/README.md).

Voltar ao [README da pasta Middleware](../README.md).
