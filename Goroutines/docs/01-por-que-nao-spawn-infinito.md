# 1. Por que não spawn infinito

## O que é uma goroutine (sem mito)

Uma **goroutine** é uma função que o runtime Go agenda para rodar **concorrentemente** com o resto do programa. O scheduler multiplexa muitas goroutines em poucas threads do sistema operacional.

Isso **não** significa:

- custo zero;
- paralelismo ilimitado em CPU;
- solução automática para performance.

## Custo real

Cada goroutine reserva **stack** (cresce sob demanda, mas existe custo) e entra na fila do scheduler. Dezenas de milhares de goroutines ativas podem:

- aumentar uso de memória;
- aumentar contenção no scheduler;
- mascarar gargalos (o problema era I/O serializado ou lock, não “falta de goroutine”).

No curso, exercícios com **50.000 goroutines** servem para ver **race** e sincronização — não é um padrão de deploy.

## Concorrência vs paralelismo

| Termo | Significado |
|--------|-------------|
| **Concorrência** | Estrutura do programa: várias tarefas em progresso (ex.: esperando rede). |
| **Paralelismo** | Execução **ao mesmo tempo** em mais de um núcleo de CPU. |

Muitas goroutines bloqueadas em `http.Get` ou query SQL são concorrência saudável **se limitadas**. Milhares de goroutines em loop de CPU sem limite raramente são.

## Quando **não** abrir outra goroutine

- Trabalho **rápido e sequencial** (soma de slice pequeno) — o overhead do `go` pode custar mais que o ganho.
- Você não sabe **como parar** a goroutine (sem `context`, sem fechar canal, sem condição de saída).
- Para cada item de uma lista enorme **sem** pool: prefira **N workers** + fila (ver doc 4).
- Para propagar erro ao `main`: use **errgroup** ou canal de erros, não `go` “solto”.

## Quando faz sentido

- Servidor HTTP (o próprio `net/http` já usa goroutine por request).
- Pool de workers processando jobs de um canal.
- Pipeline: estágio A alimenta estágio B.
- Operações I/O paralelas com **limite** (semáforo ou pool).

## Regra prática

Antes de `go func()`:

1. Quem **cancela** ou **sinaliza fim**?
2. Qual o **limite** máximo de goroutines ativas?
3. Onde o **erro** aparece se essa goroutine falhar?

Se não houver resposta clara, redesenhe antes de spawnar.

Próximo: [02-context-e-cancelamento.md](02-context-e-cancelamento.md).
