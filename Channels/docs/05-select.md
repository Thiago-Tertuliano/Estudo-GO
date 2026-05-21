# 5. Select

## O que faz

`select` espera **até um** dos casos de canal poder prosseguir (send ou receive). Se vários estiverem prontos, escolhe **pseudoaleatoriamente** entre eles.

```go
select {
case v := <-ch1:
    // ...
case ch2 <- x:
    // ...
case <-time.After(time.Second):
    // timeout
default:
    // nunca bloqueia (cuidado com busy loop)
}
```

## Timeout

Útil para não ficar preso para sempre em um receive:

```go
select {
case msg := <-ch:
    use(msg)
case <-time.After(500 * time.Millisecond):
    return errors.New("timeout")
}
```

Em serviços, prefira `context.WithTimeout` na operação de I/O quando a biblioteca suporta; `time.After` em loop quente pode vazar timers se mal usado — em loops use `time.NewTimer` e `Stop()`.

## `default`

Executa se **nenhum** case bloquearia. Padrão de polling não bloqueante — fácil virar **busy loop** consumindo CPU. Use só com motivo (ou com `time.Sleep` / ticker consciente).

## Multiplexar várias fontes

Base do fan-in e do “ouvir quit + dados”:

```go
select {
case v, ok := <-ch1:
    if !ok { ch1 = nil; continue } // nil: case ignorado no select
case v := <-ch2:
    ...
case <-ctx.Done():
    return
}
```

Atribuir **`nil`** a um canal faz o `select` ignorar esse case — padrão para desligar uma perna quando o produtor fechou o canal.

## Exemplo deste repositório

[`examples/05-select-timeout`](../examples/05-select-timeout/main.go).

```powershell
go run ./examples/05-select-timeout
```

Curso: [`aula149_select_canais.go`](../../Estudos-Realizados/Curso_Aprenda_GO/exercicios/aula149_select_canais.go).

Próximo: [06-padroes-producao.md](06-padroes-producao.md).
