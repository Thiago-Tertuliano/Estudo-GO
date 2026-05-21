# 7. Armadilhas

## Chamar `next` após 429

Se `Allow()` retornou false, **return** sem `next.ServeHTTP` — senão o handler roda mesmo após já ter respondido (ou tenta responder duas vezes).

## Mapa de IPs sem TTL

`PerIP` neste repo **nunca remove** entradas — memória cresce com IPs distintos (scan, botnet). Em produção: TTL, LRU ou Redis.

## Rate limit só na app com 10 réplicas

Cada pod tem contador próprio — limite efetivo vira `10 × configurado`. Use store **centralizado** ou limite no load balancer.

## Confiar em `X-Forwarded-For` sem validar

Atrás de proxy, o IP real pode estar em `X-Real-IP` / `X-Forwarded-For` — só confie se o proxy **sob seu controle** remove/spoof do cliente.

## `Wait` no hot path

`limiter.Wait(r.Context())` bloqueia a goroutine da request — pode esgotar threads sob carga. Para APIs públicas, prefira **`Allow()` + 429**.

## Limite muito baixo em health check

`/health` e `/ready` costumam ficar **fora** do rate limit para o orquestrador (Kubernetes).

## Próximos passos

- Rate limit por rota (`/login` mais restrito)
- Integração com [Keycloak](../../Keycloak/) (quota por client OIDC)
- Redis para cluster

---

[Índice](README.md) · [README Rate-Limiting](../README.md)
