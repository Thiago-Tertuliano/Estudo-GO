# 4. Worker pool e backpressure

## Ideia

Em vez de `go process(item)` para cada um dos 100.000 itens:

1. Um **produtor** envia **jobs** para um canal (fila).
2. **N workers** fixos leem do canal e processam.
3. Quando não houver mais jobs, o produtor **fecha** o canal; workers terminam ao esvaziar a fila.

Paralelismo máximo = **N** (mais o produtor), previsível para memória e conexões externas (DB, HTTP).

## Canal com buffer = backpressure

```go
jobs := make(chan Job, numWorkers) // buffer ≈ número de workers
```

- **Buffer 0**: cada envio bloqueia até um worker ler (sincronização forte).
- **Buffer pequeno**: suaviza picos sem acumular fila infinita na memória.
- **Sem limite conceitual**: produtor muito rápido + buffer gigante = fila enorme (evite).

Backpressure significa: quando os workers estão lentos, o **produtor desacelera** (bloqueia no `jobs <-`).

## Ciclo de vida do canal de jobs

1. `make(chan Job, buffer)`
2. Iniciar workers que fazem `for j := range jobs`
3. Produtor envia todos os jobs
4. **`close(jobs)`** — sinal de que não virá mais job
5. `wg.Wait()` nos workers

Só o **produtor** deve fechar o canal de jobs (regra: quem envia, fecha quando acabou).

## Exemplo deste repositório

[`examples/02-worker-pool`](../examples/02-worker-pool/main.go):

- 3 workers, 10 jobs;
- canal com buffer `numWorkers`;
- `close(jobs)` após o loop do produtor.

Rodar:

```powershell
go run ./examples/02-worker-pool
```

## ETL / batch (cenário 3CON)

Padrão típico:

- Ler lote de staging com `LIMIT` (offset ou cursor);
- Enviar linhas para canal de jobs;
- Pool de tamanho alinhado ao pool de conexões do banco;
- `context` cancela o lote se o job pai falhar.

Nunca “uma goroutine por linha do DW” sem teto.

## Semáforo (alternativa ao pool explícito)

Limite de goroutines em flight sem workers permanentes:

```go
sem := make(chan struct{}, 10)
sem <- struct{}{}
defer func() { <-sem }()
go func() { ... }()
```

Útil para fan-out pontual; pool fixo é mais claro para filas longas.

Próximo: [05-pipeline-e-fan-out-fan-in.md](05-pipeline-e-fan-out-fan-in.md).
