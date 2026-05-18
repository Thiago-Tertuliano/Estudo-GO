# 2. Context e cancelamento

## Papel do `context.Context`

`context` carrega:

- **cancelamento** (manual ou por timeout/deadline);
- **valores** de request (tracing, request ID — use com moderação).

Em serviços, o `context` da **request HTTP** ou do **job batch** deve ser passado para goroutines filhas e para chamadas `database/sql`, `http.Client`, gRPC, etc., para que tudo pare quando o cliente desistir ou o timeout estourar.

## API que você usa todo dia

```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()

ctx, cancel := context.WithCancel(parent)
defer cancel()

<-ctx.Done()   // canal fechado quando cancelado ou expirado
ctx.Err()      // context.Canceled ou context.DeadlineExceeded
```

**Sempre** chame `cancel()` no `defer` de `WithCancel` / `WithTimeout` / `WithDeadline` para liberar recursos do timer na árvore de contexts.

## Padrão no worker

No loop da goroutine:

```go
select {
case <-ctx.Done():
    return
default:
    // trabalho curto ou select com outros cases
}
```

Para trabalho que **bloqueia**, prefira APIs que aceitam `ctx` (`http.NewRequestWithContext`, `QueryContext`) em vez de só checar `Done()` entre sleeps.

## Exemplo deste repositório

[`examples/01-context-cancel`](../examples/01-context-cancel/main.go):

- `main` cria `context.WithTimeout` de 2 segundos;
- `worker` sai do loop quando `ctx.Done()` dispara;
- sem isso, o worker rodaria para sempre (leak lógico).

Rodar:

```powershell
go run ./examples/01-context-cancel
```

## HTTP (visão profissional)

`r.Context()` no handler é o context da request. Se o cliente fecha a conexão, o context é cancelado — queries longas devem respeitar isso.

## Erros comuns

| Erro | Consequência |
|------|----------------|
| Ignorar `ctx` em goroutine filha | Trabalho continua após timeout/cancel |
| Guardar `context` em struct global | Vazamento de cancel e comportamento indefinido |
| Passar `context.Background()` em código que deveria herdar o pai | Cancelamento nunca propaga |

Próximo: [03-waitgroup-e-errgroup.md](03-waitgroup-e-errgroup.md).
