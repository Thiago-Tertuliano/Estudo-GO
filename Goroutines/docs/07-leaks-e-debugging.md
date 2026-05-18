# 7. Leaks e debugging

## Goroutine leak

Goroutine que **nunca termina** porque ficou bloqueada em:

- leitura de canal sem remetente;
- `select` sem `case <-ctx.Done()`;
- loop infinito sem condição de saída;
- `WaitGroup` com `Add`/`Done` desbalanceados.

Sintoma em produção: memória sobe, CPU do scheduler sobe, shutdown nunca completa.

## Checklist de prevenção

- [ ] Todo `go` tem retorno garantido (cancel, `close`, fim de `range`).
- [ ] Canais fechados pelo lado que **termina de enviar**.
- [ ] `context.WithCancel` + `defer cancel()`.
- [ ] Timeouts em I/O externo.

## `go test -race`

Detector de **data race** em tempo de execução (testes ou binário com `-race`).

```powershell
cd Estudo-GO/Goroutines
go test -race ./...
```

Para exemplos que são só `main`, extraia lógica testável ou rode:

```powershell
go run -race ./examples/02-worker-pool
```

`-race` deixa o programa mais lento — use em CI e desenvolvimento, não necessariamente em binário de produção.

## pprof — goroutines

Em serviço com HTTP, importe:

```go
import _ "net/http/pprof"
// http.ListenAndServe("localhost:6060", nil)
```

Depois:

```powershell
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/goroutine
```

Procure stacks repetidas presas em `chan receive` ou `sync.Wait`.

## Logs estruturados

Em workers, logue início/fim com ID do job e `routine`/`worker_id` — facilita correlacionar goroutine presa com trabalho.

## `runtime.NumGoroutine()`

Útil em testes de integração: após shutdown, o número deve voltar perto do baseline (não é prova formal sozinho, mas ajuda).

## Debug mental rápido

1. O programa termina no `main` mas processo não morre? → goroutine órfã.
2. Shutdown trava? → `Wait` sem `Done`, ou canal bloqueado.
3. Contador errado intermitente? → race → `-race` + mutex/canal/redesign.

Próximo: [08-shutdown-gracioso.md](08-shutdown-gracioso.md).
