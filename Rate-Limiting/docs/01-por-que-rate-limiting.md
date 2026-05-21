# 1. Por que rate limiting

## O problema

Uma API HTTP aceita requisições concorrentes — em Go, cada uma costuma rodar em uma **goroutine** ([Goroutines](../../Goroutines/docs/01-por-que-nao-spawn-infinito.md)). Sem controle de taxa:

- um script com loop `curl` ou um bot pode gerar milhares de req/s;
- o processo esgota CPU, file descriptors ou pool do banco;
- usuários legítimos recebem timeout e erros em cascata.

## Motivos para limitar

| Motivo | Exemplo |
|--------|---------|
| **Proteção de recursos** | Máximo de 100 req/s na instância |
| **Fairness** | Nenhum IP consome 90% da capacidade |
| **SLA / plano comercial** | Plano free: 1000 req/dia |
| **Segurança** | Mitigar brute force em `/login` |
| **Proteção downstream** | ETL não recebe rajada ilimitada no staging |

Rate limit na **aplicação** complementa CDN, API gateway e WAF — não substitui proteção de rede.

## O que o cliente recebe

Quando o limite é excedido, a API responde em geral com:

- **HTTP 429 Too Many Requests**
- Header opcional **`Retry-After`** (segundos até tentar de novo)
- Headers informativos: `X-RateLimit-Limit`, `X-RateLimit-Remaining` (convenção de mercado)

## Onde implementar

- **Edge** (Cloudflare, Kong, NGINX): antes da app.
- **Middleware na app** (esta trilha): `func(http.Handler) http.Handler` — ver [Middleware](../../Middleware/docs/02-assinatura-middleware.md).

Próximo: [02-conceitos-e-algoritmos.md](02-conceitos-e-algoritmos.md).
