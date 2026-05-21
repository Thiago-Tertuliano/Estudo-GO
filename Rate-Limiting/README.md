# Rate Limiting — estudo profissional em Go

Material sobre **limitação de taxa** em APIs HTTP: por que existe, algoritmos (token bucket, janela fixa), implementação com `golang.org/x/time/rate` e middleware que responde **429 Too Many Requests**.

## Pré-requisitos

- Go 1.22+
- Trilha [Middleware](../Middleware/) (assinatura `func(http.Handler) http.Handler` e chain)
- HTTP básico

## Por que estudar isso

Sem rate limit, um cliente ou bot pode saturar CPU, conexões de banco e goroutines do servidor. O limite protege o serviço, garante fairness entre clientes e sustenta SLAs por plano.

## Conteúdo

| Item | Descrição |
|------|-----------|
| [`docs/`](docs/) | 7 capítulos em Markdown |
| [`examples/`](examples/) | 7 programas executáveis |
| [`internal/ratelimit/`](internal/ratelimit/) | Middleware global e por IP |

### Documentação

Índice: [`docs/README.md`](docs/README.md).

### Exemplos

| # | Pasta | Tema |
|---|-------|------|
| 01 | `01-token-bucket-basico` | `rate.Limiter` sem HTTP |
| 02 | `02-middleware-global` | Um limiter para toda a API |
| 03 | `03-middleware-per-ip` | Limiter por IP |
| 04 | `04-fixed-window-manual` | Janela fixa com mutex (didático) |
| 05 | `05-stdlib-server` | Servidor com chain + rate limit |
| 06 | `06-chi-bridge` | Chi + `ratelimit.Middleware` |
| 07 | `07-echo-bridge` | Echo + adaptador stdlib |

## Ordem de estudo

1. Docs `01`–`03` + exemplo `01`
2. Docs `04`–`05` + exemplos `02`–`04`
3. Exemplo `05` + doc `07`
4. Doc `06` + exemplos `06`–`07`

## Comandos

```powershell
cd Estudo-GO\Rate-Limiting
go mod tidy
go build ./...

go run ./examples/01-token-bucket-basico
go run ./examples/05-stdlib-server

# Disparar várias requisições até ver 429:
1..20 | ForEach-Object { curl.exe -s -o NUL -w "%{http_code}\n" http://localhost:8080/ }
```

## Fora do escopo desta trilha

- Rate limit distribuído (Redis/Valkey) em cluster
- gRPC (apenas menção nos docs)
- Testes de carga com k6

## Referências

- [golang.org/x/time/rate](https://pkg.go.dev/golang.org/x/time/rate)
- [RFC 6585 — 429 Too Many Requests](https://www.rfc-editor.org/rfc/rfc6585)

---

Voltar ao índice geral: [`../README.md`](../README.md).
