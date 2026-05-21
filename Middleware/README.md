# Middleware HTTP — estudo profissional em Go

Material focado em **como funciona** middleware HTTP na stdlib (`net/http`) e **como construir do zero**, antes de usar pacotes prontos do Chi ou Echo.

## Objetivo

Entender o tipo `func(http.Handler) http.Handler`, a ordem da cadeia (onion), wrappers de `ResponseWriter`, valores em `context` e recover de panic — e reconhecer o mesmo padrão em `r.Use(...)` e `e.Use(...)`.

## Pré-requisitos

- Go 1.22+
- HTTP básico (método, path, status, headers)
- Familiaridade com [Goroutines/context](../Goroutines/docs/02-context-e-cancelamento.md) (cada request roda em uma goroutine)

## Relação com seus projetos

| Projeto | Framework | Middleware usado |
|---------|-----------|------------------|
| [Curso-API-GO](../Estudos-Realizados/Curso-API-GO/main.go) | Chi | Logger, Recoverer, RequestID, Timeout |
| [API-Kelche](../Estudos-Realizados/API-Kelche/fitness-api/main.go) | Echo | Logger, CORS |

## Conteúdo

| Item | Descrição |
|------|-----------|
| [`docs/`](docs/) | 7 capítulos em Markdown |
| [`examples/`](examples/) | 10 programas (`go run ./examples/...`) |
| [`internal/httpmw/`](internal/httpmw/) | Middlewares reutilizáveis nos exemplos 03+ |

### Documentação

Índice: [`docs/README.md`](docs/README.md).

### Exemplos

| # | Pasta | Tema |
|---|-------|------|
| 01 | `01-hello-handler` | `HandlerFunc` mínimo |
| 02 | `02-wrap-handler` | Primeiro wrap manual |
| 03 | `03-logging` | Log método, path, duração |
| 04 | `04-status-recorder` | Capturar status HTTP |
| 05 | `05-chain` | Empilhar middlewares |
| 06 | `06-context-value` | Request ID no context |
| 07 | `07-recover-panic` | Recover de panic |
| 08 | `08-stdlib-server` | Servidor completo (mux + chain) |
| 09 | `09-chi-bridge` | Mesmo middleware + `chi.Use` |
| 10 | `10-echo-bridge` | Adaptador para Echo |

## Ordem de estudo

1. Docs `01`–`03` + exemplos `01`–`05`
2. Docs `04`–`05` + exemplos `06`–`07`
3. Exemplo `08` + doc `03`/`07`
4. Doc `06` + exemplos `09`–`10`

## Comandos

```powershell
cd Estudo-GO\Middleware
go mod tidy
go build ./...

# Servidor (Ctrl+C para parar)
go run ./examples/08-stdlib-server
go run ./examples/09-chi-bridge
go run ./examples/10-echo-bridge

# Exemplos pontuais (sem servidor longo)
go run ./examples/01-hello-handler
```

Teste com curl (porta padrão `:8080` nos servidores):

```powershell
curl http://localhost:8080/
curl http://localhost:8080/panic
```

## Referências

- [net/http — Handlers](https://pkg.go.dev/net/http#Handler)
- [Chi middleware](https://github.com/go-chi/chi#middlewares)
- [Echo middleware](https://echo.labstack.com/docs/middleware)

---

Voltar ao índice geral: [`../README.md`](../README.md).
