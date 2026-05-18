# 5. Pipeline e fan-out / fan-in

## Pipeline

Cadeia de estágios conectados por **canais**:

```
gerar → processar → agregar
  |         |            |
 raw      mid          (soma)
```

Cada estágio:

- lê de um canal de entrada;
- escreve em um canal de saída;
- **fecha** o canal de saída quando termina (para o próximo estágio saber que acabou).

O estágio final costuma ser síncrono (`range` até fechar) ou outra goroutine com `WaitGroup`.

## Exemplo deste repositório

[`examples/04-pipeline`](../examples/04-pipeline/main.go):

- `generate` envia inteiros e fecha `raw`;
- `process` dobra valores em `mid` e fecha `mid`;
- `aggregate` soma tudo após `mid` fechar.

Rodar:

```powershell
go run ./examples/04-pipeline
```

## Fan-out

**Vários** workers leem do **mesmo** canal de entrada (competição justa pelo próximo item):

```go
for w := 0; w < n; w++ {
    go func() {
        for job := range jobs {
            // processar
        }
    }()
}
```

É o mesmo padrão do worker pool: fan-out do trabalho.

## Fan-in

**Vários** produtores enviam para **um** canal de saída (ou um estágio agrega resultados):

```go
// várias goroutines fazem: results <- valor
// uma goroutine: for v := range results { ... }
```

Cuidado: fechar `results` só quando **todos** os produtores terminarem — use `WaitGroup` + uma goroutine que fecha após `wg.Wait()`, ou `errgroup`.

## Backpressure no pipeline

Buffers pequenos entre estágios (`make(chan T, 2)`) fazem o estágio lento **segurar** o rápido, evitando buffer infinito na memória.

## Quando pipeline vs pool simples

| Situação | Preferência |
|----------|-------------|
| Um tipo de trabalho homogêneo | Worker pool |
| Estágios diferentes (parse → validate → persist) | Pipeline |
| Muitos produtores, um consumidor | Fan-in com merge controlado |

Próximo: [06-mutex-atomic-e-canais.md](06-mutex-atomic-e-canais.md).
