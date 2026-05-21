# 4. Canais direcionais

## Sintaxe

| Tipo | Quem pode |
|------|-----------|
| `chan T` | Send e receive (bidirecional) |
| `chan<- T` | **Só send** |
| `<-chan T` | **Só receive** |

Conversão implícita: um `chan T` pode ser passado onde a função espera `chan<- T` ou `<-chan T` (perde-se um lado da operação no tipo).

## Por que usar em APIs

Deixa explícito no **contrato da função**:

- `gen(...) <-chan int` — retorna só leitura; quem chama não deve enviar nesse canal.
- `consume(in <-chan int)` — só recebe; não pode enviar por engano.

Reduz bugs em pipelines grandes.

## Exemplo: gen + transform

[`examples/04-directional`](../examples/04-directional/main.go):

- `gen` retorna `<-chan int` e fecha o canal internamente;
- `square` recebe `<-chan int`, retorna `<-chan int`;
- `main` só faz `range` no resultado.

```powershell
go run ./examples/04-directional
```

## Curso

Padrão parecido com [`aula154_exercicio4_select_quit.go`](../../Estudos-Realizados/Curso_Aprenda_GO/exercicios/aula154_exercicio4_select_quit.go) (`chan<- int`, `<-chan int`).

## Erro comum

Assinar `done chan<- struct{}` e dentro da função fazer `<-done` — **compile error**: não se recebe de canal send-only. Quem **escuta** parada usa `<-chan struct{}`:

```go
func worker(done <-chan struct{}) { ... case <-done: ... }
```

Ver também [06-select-quit](../examples/06-select-quit/main.go).

Próximo: [05-select.md](05-select.md).
