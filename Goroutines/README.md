# Goroutines — estudo profissional em Go

Material focado em **concorrência aplicada**: quando usar goroutines, como encerrar trabalho, limitar paralelismo e evitar bugs que só aparecem em produção (leak, race, cancelamento ignorado).

Não é um repositório de “disparar mil goroutines e imprimir contador”. O objetivo é o mesmo critério de serviços reais: **APIs**, **workers**, **batch/ETL** e integrações com I/O.

## Para quem já estudou o básico

No curso em [`../Estudos-Realizados/Curso_Aprenda_GO/`](../Estudos-Realizados/Curso_Aprenda_GO/) você já viu:

- `aula126` — WaitGroup
- `aula139` — race, mutex, atomic
- `aula146+` — canais, select, context

Aqui a progressão é **uso em sistema**: lifecycle, pools, pipelines, shutdown e ferramentas (`-race`, `pprof`).

## Acadêmico vs profissional (resumo)

| Abordagem acadêmica | Abordagem deste estudo |
|---------------------|-------------------------|
| Goroutine como “thread leve” | Goroutine com **owner**, **cancelamento** e **limite** de concorrência |
| 50k goroutines no contador | Worker pool, backpressure, `context` |
| Mutex só no exercício de race | Estado compartilhado **mínimo**; preferir fluxo por canais |
| `go func()` sem tratar erro | Erro propagado (`errgroup`, canal de erros, logs) |

## Conteúdo

| Item | Descrição |
|------|-----------|
| [`docs/`](docs/) | Notas em Markdown (conceitos, padrões, armadilhas) |
| [`examples/`](examples/) | Programas pequenos por tema (`go run ./examples/...`) |

### Documentação (`docs/`)

Índice: [`docs/README.md`](docs/README.md).

1. [Por que não spawn infinito](docs/01-por-que-nao-spawn-infinito.md)
2. [Context e cancelamento](docs/02-context-e-cancelamento.md)
3. [WaitGroup e errgroup](docs/03-waitgroup-e-errgroup.md)
4. [Worker pool e backpressure](docs/04-worker-pool-e-backpressure.md)
5. [Pipeline e fan-out/fan-in](docs/05-pipeline-e-fan-out-fan-in.md)
6. [Mutex, atomic e canais](docs/06-mutex-atomic-e-canais.md)
7. [Leaks e debugging](docs/07-leaks-e-debugging.md)
8. [Shutdown gracioso](docs/08-shutdown-gracioso.md)

### Exemplos (`examples/`)

| Exemplo | Ideia central |
|---------|----------------|
| `01-context-cancel` | Worker para quando o `context` cancela |
| `02-worker-pool` | Fila de jobs + pool limitado |
| `03-errgroup` | Paralelismo com propagação de erro |
| `04-pipeline` | Estágios com backpressure |
| `05-graceful-shutdown` | Cancel + WaitGroup no encerramento |

## Regras práticas (checklist mental)

- [ ] Toda goroutine tem **caminho de saída** (`ctx.Done()`, canal fechado, fim do loop).
- [ ] I/O externo (HTTP, DB, fila) usa **`context.Context`** da operação pai.
- [ ] Paralelismo **limitado** (pool ou semáforo), não uma goroutine por item sem teto.
- [ ] Erros dentro de `go` **não** são engolidos (log + retorno ao grupo principal).
- [ ] CI ou hábito local: `go test -race ./...` nos pacotes com concorrência.

## Comandos úteis

```powershell
go run ./examples/01-context-cancel
# ou: cd examples\<nome-do-exemplo> && go run .

# Detectar data race
go test -race ./...

# Ver goroutines (com exemplo que exponha :6060/debug/pprof quando aplicável)
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/goroutine
```

## Onde isso aparece no dia a dia

- **HTTP (`net/http`)**: cada request já roda em goroutine; o profissional trabalha com **middleware**, **timeout** e **context** da request.
- **Batch / ETL**: workers processando lotes com limite de memória e conexões ao banco.
- **Consumers**: ler fila com N workers, commit/ack após processar, shutdown sem perder mensagem (conceito; implementação depende do broker).

## Referências

- [Effective Go — Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go blog — Share memory by communicating](https://go.dev/blog/codelab-share)
- [Package context](https://pkg.go.dev/context)
- [x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [Go memory model](https://go.dev/ref/mem)

## Pré-requisitos

- Go 1.21+ (ou a versão do seu `go.mod` no monorepo de estudos)
- Familiaridade com funções, structs e interfaces (nível do `Curso_Aprenda_GO`)

---

Voltar ao índice geral: [`../README.md`](../README.md).
