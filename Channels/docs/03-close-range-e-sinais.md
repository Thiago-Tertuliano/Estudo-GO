# 3. Close, range e sinais

## Fechar um canal

Só quem **envia** deve chamar `close(ch)` (quando não houver mais valores a enviar).

```go
close(ch)
```

Efeitos:

- Receives subsequentes obtêm o valor zero de `T` e `ok == false`:

  ```go
  v, ok := <-ch
  if !ok { /* canal fechado e esvaziado */ }
  ```

- `for v := range ch` termina quando o canal está fechado **e** vazio.

## O que causa panic

- **Send** em canal já fechado: `panic: send on closed channel`.
- **Close** de canal já fechado: `panic: close of closed channel`.
- **Close** do lado que só recebe — conceitualmente errado; quem fecha é o produtor.

## Padrão produtor / consumidor

```go
func producer(out chan<- int) {
    defer close(out)
    for i := 0; i < n; i++ {
        out <- i
    }
}

func main() {
    ch := make(chan int)
    go producer(ch)
    for v := range ch {
        fmt.Println(v)
    }
}
```

## Canal `struct{}` como sinal

`chan struct{}` sem payload — só evento (fechado, ou um send vazio):

```go
done := make(chan struct{})
close(done)           // broadcast “pare” para quem faz <-done
// ou: done <- struct{}{}
```

Em código novo, muitas vezes **`context.Context`** substitui canal de quit; canais de sinal ainda aparecem em código legado e em padrões didáticos.

## Exemplo deste repositório

[`examples/03-close-range`](../examples/03-close-range/main.go).

```powershell
go run ./examples/03-close-range
```

Próximo: [04-canais-direcionais.md](04-canais-direcionais.md).
