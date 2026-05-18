# 6. Mutex, atomic e canais

## Princípio do Go

> *Don't communicate by sharing memory; share memory by communicating.*

Na prática: prefira passar dados por **canais** ou confinar mutação em **um** goroutine dono do estado. Use **mutex** ou **atomic** quando o modelo de canais ficar artificial ou for hot path medido.

## Race condition

Duas goroutines acessam a mesma variável, **pelo menos uma escreve**, sem sincronização → **data race** (comportamento indefinido).

No curso: [`aula139_exercicio3_goroutines_race.go`](../../Estudos-Realizados/Curso_Aprenda_GO/exercicios/aula139_exercicio3_goroutines_race.go) — contador incorreto sem proteção.

Detectar:

```powershell
go test -race ./...
go run -race ./examples/02-worker-pool
```

## `sync.Mutex`

Protege seção crítica:

```go
var mu sync.Mutex
var total int

mu.Lock()
total++
mu.Unlock()
```

- Use **`defer mu.Unlock()`** após `Lock` em funções com vários returns.
- Encapsule contador em struct com métodos `Add()` que trancam internamente — API menor e mais segura.

Versão do curso: `aula139_exercicio4_goroutines_mutex.go`.

## `sync/atomic`

Operações atôicas em inteiros (`AddInt64`, `Load`, `Store`). Útil para métricas e contadores simples.

Versão do curso: `aula139_exercicio5_goroutines_atomic.go`.

**Não** substitui mutex para estruturas complexas ou múltiplos campos que precisam ficar consistentes juntos.

## Canais

| Uso | Canal |
|-----|--------|
| Fila de trabalho | `chan Job` + workers |
| Sinalizar fim | `close(ch)` + `range` |
| Cancelamento | `context.Done()` (preferível a canal custom para cancel) |
| Resultado único | `chan Result` com buffer 1 ou `select` |

## Como escolher

| Cenário | Sugestão |
|---------|----------|
| Fila de jobs, pipeline | Canal |
| Contador global em hot path | `atomic` ou métrica do Prometheus |
| Cache/map compartilhado | `sync.RWMutex` ou `sync.Map` (casos específicos) |
| Estado só de um worker | Variável local, sem lock |

## Exemplos deste repo

`02-worker-pool` e `04-pipeline` evitam estado global mutável — cada job flui pelos canais.

Próximo: [07-leaks-e-debugging.md](07-leaks-e-debugging.md).
