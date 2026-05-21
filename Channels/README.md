# Channels — estudo profissional em Go

Material focado em **canais** como mecanismo de comunicação entre goroutines: buffer, `close`, direção (`chan<-` / `<-chan`), `select`, fan-in e armadilhas de produção.

Complementa a trilha [Goroutines/](../Goroutines/) (lifecycle, pools, context). Lá os canais aparecem como **fila**; aqui a semântica do canal é o centro.

## Para quem já estudou o básico

No curso em [`../Estudos-Realizados/Curso_Aprenda_GO/`](../Estudos-Realizados/Curso_Aprenda_GO/):

- `aula146` — send/receive
- `aula149` — `select` em dois canais
- `aula154` — canais direcionais + quit

## Acadêmico vs profissional

| Abordagem acadêmica | Abordagem deste estudo |
|---------------------|-------------------------|
| Canal sem buffer “para tudo” | Buffer dimensionado (backpressure) |
| `close` em qualquer lugar | Só o **produtor** fecha |
| Quit channel sempre | Preferir **`context`** quando possível |
| `select` com `default` em loop | Evitar busy loop; timeout com critério |

## Conteúdo

| Item | Descrição |
|------|-----------|
| [`docs/`](docs/) | 7 capítulos em Markdown |
| [`examples/`](examples/) | Programas `01` … `07` |

### Documentação (`docs/`)

Índice: [`docs/README.md`](docs/README.md).

1. [O que é um canal](docs/01-o-que-e-canal.md)
2. [Buffered vs unbuffered](docs/02-buffered-vs-unbuffered.md)
3. [Close, range e sinais](docs/03-close-range-e-sinais.md)
4. [Canais direcionais](docs/04-canais-direcionais.md)
5. [Select](docs/05-select.md)
6. [Padrões em produção](docs/06-padroes-producao.md)
7. [Armadilhas e leaks](docs/07-armadilhas-e-leaks.md)

### Exemplos (`examples/`)

| Exemplo | Ideia central |
|---------|----------------|
| `01-basics` | Canal unbuffered (handshake) |
| `02-buffered` | Capacidade e bloqueio quando cheio |
| `03-close-range` | `close` + `for range` |
| `04-directional` | `gen` / `square` com tipos só-send / só-recv |
| `05-select-timeout` | `select` + `time.After` |
| `06-select-quit` | Parada com `<-chan struct{}` (`done` receive-only em `gen`) |
| `07-fan-in` | Dois produtores + `merge` |

## Comandos

```powershell
cd Estudo-GO\Channels
go run ./examples/01-basics
go run ./examples/02-buffered
go run ./examples/03-close-range
go run ./examples/04-directional
go run ./examples/05-select-timeout
go run ./examples/06-select-quit
go run ./examples/07-fan-in

go test -race ./...   # quando houver testes
```

## Ordem sugerida

1. Docs `01` → `05` + exemplos `01` → `05`
2. [Goroutines](../Goroutines/) worker pool e pipeline (revisitar canais na prática)
3. Docs `06`–`07` + exemplos `06`–`07`

## Referências

- [Effective Go — Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go blog — Share memory by communicating](https://go.dev/blog/codelab-share)
- [Package builtin — close](https://pkg.go.dev/builtin#close)

---

Voltar ao índice geral: [`../README.md`](../README.md).
