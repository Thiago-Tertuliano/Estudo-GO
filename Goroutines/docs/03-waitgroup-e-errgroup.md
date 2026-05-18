# 3. WaitGroup e errgroup

## `sync.WaitGroup` — esperar um lote terminar

Use quando várias goroutines fazem trabalho e o `main` (ou um estágio) precisa **bloquear até todas terminarem**, sem necessariamente propagar o **primeiro erro** de forma automática.

```go
var wg sync.WaitGroup
for i := 0; i < n; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // trabalho
    }()
}
wg.Wait()
```

### Armadilhas

- **`Add` depois do `go`** sem cuidado: race se o `Wait` rodar antes do `Add` — faça `Add` **antes** de `go`, ou use um padrão com canal de “pronto”.
- **`Add` com valor errado**: `Wait` bloqueia para sempre ou retorna cedo demais.
- **Copiar WaitGroup** para outra struct: passe **ponteiro** (`*sync.WaitGroup`).

No curso: `aula126` — WaitGroup com tarefas que simulam `Sleep`.

## `errgroup` — paralelismo com erro e cancelamento

Pacote: `golang.org/x/sync/errgroup`.

```go
g, ctx := errgroup.WithContext(parent)
g.Go(func() error { return taskA(ctx) })
g.Go(func() error { return taskB(ctx) })
if err := g.Wait(); err != nil {
    // primeira falha; ctx do grupo é cancelado
}
```

Quando **uma** goroutine retorna erro não-nil, o **context derivado** é cancelado — as outras tarefas devem observar `ctx.Done()` e sair.

### Quando preferir errgroup em vez de só WaitGroup

- Várias chamadas independentes (validar 3 APIs, buscar 3 recursos).
- Uma falha deve **abortar** o restante.
- Você já propaga `error` em todo o fluxo.

## Exemplo deste repositório

[`examples/03-errgroup`](../examples/03-errgroup/main.go):

- três `task`s em paralelo;
- a task 2 falha de propósito;
- as outras recebem cancelamento via `ctx` do grupo.

Rodar:

```powershell
go run ./examples/03-errgroup
```

## Comparativo rápido

| Ferramenta | Espera fim | Cancela irmãs no erro | Retorna `error` |
|------------|------------|------------------------|-----------------|
| WaitGroup | Sim | Não (manual) | Não |
| errgroup | Sim | Sim (com `WithContext`) | Sim (`Wait()`) |

## Erro dentro de `go` sem errgroup

Se você usar só `go func() { ... }()`:

- o `main` pode terminar antes da goroutine;
- erros ficam invisíveis;
- panics podem derrubar o processo.

Profissionalmente: **errgroup**, canal `chan error` com buffer 1, ou worker pool que agrega erros.

Próximo: [04-worker-pool-e-backpressure.md](04-worker-pool-e-backpressure.md).
