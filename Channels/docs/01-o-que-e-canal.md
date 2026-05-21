# 1. O que é um canal

## Ideia em uma frase

Um **canal** (`chan T`) é um tipo que permite **passar valores** entre goroutines com **sincronização** definida pela linguagem — a alternativa idiomática a locks para muitos fluxos de dados.

## Send e receive

```go
ch := make(chan int)   // unbuffered
ch <- 42               // send (bloqueia até alguém receber, se unbuffered)
v := <-ch              // receive
```

- **Send** e **receive** são operações que o runtime coordena.
- Em canal **sem buffer**, send e receive formam um **handshake**: as duas pontas se encontram no mesmo instante lógico.

## Por que existem

Goroutines sozinhas não compartilham memória de forma segura sem disciplina. Canais implementam o lema do Go:

> *Don't communicate by sharing memory; share memory by communicating.*

Na prática profissional: fila de jobs, resultados de pipeline, sinal de “pronto” ou “pare”.

## Canal unbuffered = sincronização forte

O produtor **espera** o consumidor estar pronto para receber (e vice-versa). Útil quando você quer **backpressure imediata** ou garantir que o próximo passo já está ouvindo.

## Exemplo deste repositório

[`examples/01-basics`](../examples/01-basics/main.go): uma goroutine envia `string`, o `main` recebe e imprime.

```powershell
go run ./examples/01-basics
```

## Relação com o curso

[`aula146_canais.go`](../../Estudos-Realizados/Curso_Aprenda_GO/exercicios/aula146_canais.go) — primeiro contato com `ch <-` e `<-ch`.

## O que vem depois

Buffer (capacidade > 0), `close`, direção explícita (`chan<-`, `<-chan`) e `select`.

Próximo: [02-buffered-vs-unbuffered.md](02-buffered-vs-unbuffered.md).
