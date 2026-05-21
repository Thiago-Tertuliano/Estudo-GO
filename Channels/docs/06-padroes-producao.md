# 6. Padrões em produção

## Worker pool (fila + N consumidores)

Canal de jobs + goroutines fixas:

```text
produtor → jobs (chan) → worker 1..N → (resultados opcional)
```

Implementação completa em [Goroutines/examples/02-worker-pool](../../Goroutines/examples/02-worker-pool/main.go). O canal aqui é a **fila**; o limite de paralelismo é o **número de workers**, não o tamanho infinito da fila.

## Pipeline

Estágios encadeados; cada estágio lê de `<-chan` e escreve em outro canal, fechando ao terminar.

- [Goroutines/examples/04-pipeline](../../Goroutines/examples/04-pipeline/main.go)
- [Channels/examples/04-directional](../examples/04-directional/main.go) — API com canais direcionais

## Fan-out / fan-in

- **Fan-out:** vários workers leem do mesmo canal de jobs.
- **Fan-in:** várias fontes convergem para um consumidor (merge ou `select` em loop).

[`examples/07-fan-in`](../examples/07-fan-in/main.go):

- dois `producer` em canais `a` e `b`;
- após `WaitGroup`, `close(a)` e `close(b)`;
- `merge` multiplexa para `merged`; `main` faz `range merged`.

```powershell
go run ./examples/07-fan-in
```

Versão didática com dois canais e `select` fixo: curso `aula149`.

## Select + quit (parada cooperativa)

[`examples/06-select-quit`](../examples/06-select-quit/main.go):

- `done <-chan struct{}` — **receive-only** em `gen` (quem escuta parada);
- `main` faz `defer close(done)`;
- dentro do loop, `select` entre enviar em `out` e `<-done`.

```powershell
go run ./examples/06-select-quit
```

Em código novo: combine com **`context.Context`** (`case <-ctx.Done()`) em vez de reinventar quit channel, salvo integração com APIs legadas.

## Semáforo com canal

Limitar goroutines em flight:

```go
sem := make(chan struct{}, 10)
sem <- struct{}{}        // adquire
defer func() { <-sem }() // libera
```

Equivalente conceitual a “no máximo 10 operações paralelas”.

## Quando **não** usar canal

- Contador ou cache compartilhado simples → `sync.Mutex` ou tipo com mutex interno.
- “Só esperar N goroutines” → `sync.WaitGroup` ou `errgroup`.
- Cancelamento hierárquico → `context`.

Canal é para **fluxo de dados** e sincronização desse fluxo, não substituto universal de lock.

Próximo: [07-armadilhas-e-leaks.md](07-armadilhas-e-leaks.md).
