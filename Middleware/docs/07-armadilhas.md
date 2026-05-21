# 7. Armadilhas

## Esquecer `next.ServeHTTP`

Middleware que só faz trabalho "antes" e não chama `next` — handler nunca roda.

## Responder duas vezes

Chamar `http.Error` ou `Write` no middleware **e** deixar o handler escrever de novo — cliente recebe body inválido. Middleware de auth deve **retornar** sem chamar `next` se negar.

## Panic sem recover

Panic no handler derruba o processo inteiro se não houver middleware de recover **por fora**. Exemplo: [`07-recover-panic`](../examples/07-recover-panic/main.go).

## Timeout

`middleware.Timeout` (Chi) cancela `context` — handlers devem checar `r.Context().Err()` em operações longas.

## Context values

- Chaves string genéricas colidem.
- Guardar ponteiros mutáveis no context — race.

## Logging após `WriteHeader` errado

Se outro middleware já escreveu resposta, seu log pode ver status 200 no recorder mas cliente já recebeu outro fluxo — mantenha ordem clara.

## Próximo passo (fora desta trilha)

- Middleware de autenticação (Bearer JWT / sessão)
- Integração com [Keycloak](../../Keycloak/) ou validação JWT local
- Testes com `httptest.ResponseRecorder` + middleware chain

---

[Índice](README.md) · [README Middleware](../README.md)
