# 8. Shutdown gracioso

## Objetivo

Ao receber **SIGINT** / **SIGTERM** (Kubernetes, systemd, Ctrl+C no terminal):

1. Parar de aceitar **novo** trabalho (ou cancelar context).
2. Deixar workers **terminar** o que já começaram (dentro de um timeout).
3. Fechar conexões (DB, HTTP server `Shutdown`).

Evita cortar no meio sem flush e reduz goroutines órfãs.

## Padrão usado neste estudo

1. `signal.Notify` para `os.Interrupt` e `syscall.SIGTERM`.
2. `context.WithCancel` compartilhado pelos workers.
3. Ao sinal: `cancel()`.
4. Workers observam `<-ctx.Done()` e saem do loop.
5. `sync.WaitGroup` espera todos; opcional **timeout** se algum travar.

## Exemplo deste repositório

[`examples/05-graceful-shutdown`](../examples/05-graceful-shutdown/main.go):

- 3 workers imprimindo periodicamente;
- Ctrl+C → cancel → `wg.Wait()` com timeout de 3s no `main`.

Rodar:

```powershell
go run ./examples/05-graceful-shutdown
```

Pressione **Ctrl+C** e observe as mensagens `shutdown`.

## `http.Server` (produção)

```go
srv := &http.Server{Addr: ":8080", Handler: mux}
go func() { _ = srv.ListenAndServe() }()

<-sigCh
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
_ = srv.Shutdown(ctx) // para de aceitar; espera requests ativas
```

`Shutdown` é o equivalente HTTP do padrão cancel + wait dos workers.

## Ordem típica em serviço

1. Cancelar context raiz do app.
2. `WaitGroup` nos workers / consumers.
3. `Shutdown` no servidor HTTP.
4. Fechar pool de DB (`db.Close()`).

Timeout global: se passar de N segundos, logar e sair com código não-zero (orquestrador sobe nova instância).

## Anti-padrões

- `os.Exit(0)` imediato no handler de sinal sem esperar workers.
- Ignorar `ctx` em query longa após cancel.
- Fechar canal de jobs enquanto produtor ainda envia (panic).

---

Trilha concluída. Voltar ao [índice](README.md) ou ao [README principal](../README.md).
