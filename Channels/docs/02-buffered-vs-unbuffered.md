# 2. Buffered vs unbuffered

## Criar canal com capacidade

```go
ch := make(chan int)    // capacidade 0 — unbuffered
ch := make(chan int, 3) // buffer de 3 elementos
```

`len(ch)` = quantos valores estão no buffer agora.  
`cap(ch)` = capacidade máxima do buffer.

## Comportamento

| Tipo | Send bloqueia quando… | Receive bloqueia quando… |
|------|------------------------|-------------------------|
| **Unbuffered** | Não há receptor pronto | Não há remetente pronto |
| **Buffered** | Buffer cheio (`len == cap`) | Buffer vazio (`len == 0`) |

## Backpressure

**Backpressure** = produtor desacelera quando o consumidor não acompanha.

- **Unbuffered:** cada item força sincronização par a par.
- **Buffered pequeno:** absorve picos curtos; quando enche, o produtor bloqueia no `ch <-`.
- **Buffer enorme “por precaução”:** esconde gargalo e pode estourar memória — evite em produção sem motivo.

Regra prática: tamanho do buffer ≈ número de workers ou janela de prefetch que você aceita em memória (ex.: `make(chan Job, numWorkers)` no worker pool em [Goroutines](../../Goroutines/docs/04-worker-pool-e-backpressure.md)).

## Quando usar cada um

- **Unbuffered:** handshake, “só continua quando o outro lado pegou”, sinais pontuais.
- **Buffered:** filas de trabalho, amortizar latência entre estágios de pipeline, semáforo (`make(chan struct{}, N)`).

## Exemplo deste repositório

[`examples/02-buffered`](../examples/02-buffered/main.go): três envios cabem no buffer; o quarto bloquearia até haver leitura.

```powershell
go run ./examples/02-buffered
```

Próximo: [03-close-range-e-sinais.md](03-close-range-e-sinais.md).
