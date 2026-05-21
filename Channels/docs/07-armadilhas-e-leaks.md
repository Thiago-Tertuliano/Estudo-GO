# 7. Armadilhas e leaks

## Goroutine bloqueada em send

```go
ch := make(chan int) // unbuffered, sem receptor
go func() { ch <- 1 }() // bloqueia para sempre → leak
```

Causas comuns:

- consumidor terminou e ninguém mais lê;
- `break` no `range` sem cancelar produtor;
- buffer cheio e produtor continua enviando.

**Prevenção:** `close` + `range`, `ctx.Done()` no `select`, ou garantir receptor até o fim.

## Fechar do lado errado

- Fechar canal que outros ainda **enviam** → panic no próximo send.
- Fechar duas vezes → panic.
- Consumidor fecha canal de jobs → produtor pode panicar ao enviar.

**Regra:** quem **termina de produzir** fecha o canal de saída.

## Receive sem fim

Loop infinito em `<-ch` sem o produtor fechar nunca termina. Combine com `ok` ou `range` após `close`.

## Nil channel no select

```go
var ch chan int // nil
select {
case <-ch: // nunca selecionado
}
```

Útil **de propósito** para desabilitar um case após fechar um dos inputs no fan-in (ver doc 5).

## Misturar quit channel e context

Dois mecanismos de cancelamento sem documentação confunde a equipe. Padronize: **`context` na árvore de operação**; canal de quit só se necessário para biblioteca antiga.

## `select` + `default` em loop

Gira CPU a 100% sem trabalho útil. Prefira bloquear em canal ou usar `time.Ticker`.

## Ferramentas

```powershell
go test -race ./...
go run -race ./examples/07-fan-in
```

Race em canal costuma envolver **fechar enquanto outra goroutine envia** ou variável compartilhada **além** do canal.

Para goroutines presas: [Goroutines/docs/07-leaks-e-debugging.md](../../Goroutines/docs/07-leaks-e-debugging.md) (`pprof`, checklist).

---

Trilha concluída. [Índice](README.md) · [README Channels](../README.md).
