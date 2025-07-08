# Estudos de Go

Este repositório contém uma coletânea completa de exercícios, exemplos e testes em Go, organizados por temas para facilitar o aprendizado e a consulta.

## Como usar

- Todos os exercícios estão na pasta `exercicios/`.
- Para rodar um exercício, use:
  ```bash
  go run exercicios/nome_do_arquivo.go
  ```
- Para rodar todos os testes:
  ```bash
  go test ./exercicios
  ```

## Índice Temático dos Exercícios

### 1. Conceitos Básicos
- Hello World, variáveis, tipos, valores zero, constantes, iota, operadores, controle de fluxo, loops, condicionais, switch, operadores lógicos

### 2. Arrays, Slices e Maps
- Arrays, slices (fatiamento, make, multidimensionais, append, slicing), maps, operações com maps

### 3. Structs e Métodos
- Structs simples, embutidos, anônimos, métodos, interfaces, polimorfismo

### 4. Funções e Closures
- Funções básicas, variádicas, anônimas, closure, recursividade, defer

### 5. Ponteiros
- Ponteiros, ponteiros em structs, exemplos práticos

### 6. JSON e Serialização
- Marshal, Unmarshal, Encoder, Decoder, ordenação, customização

### 7. Concorrência
- Goroutines, WaitGroups, Mutex, Atomic, canais, select, context, workers, convergência/divergência

### 8. Tratamento de Erros
- Erros, recover, erros customizados, log, exemplos práticos

### 9. Testes, Benchmarks e Ferramentas
- Testes unitários, benchmarks, testes em tabela, go fmt, go vet, golint, exemplos de documentação

---

## Exemplos de arquivos (por tema)

- **Básico:** `aula87_funcoes.go`, `aula89_defer.go`, `aula90_metodos.go`, `aula91_interfaces_polimorfismo.go`
- **Arrays/Slices/Maps:** `aula69_exercicio1_array.go`, `aula64_slice_make.go`, `aula69_exercicio8_maps.go`
- **Structs:** `aula79_struct.go`, `aula80_structs_embutidos.go`, `exercicio_struct_anonimo.go`
- **Funções/Closures:** `aula96_closure.go`, `aula97_recursividade.go`, `aula98_exercicio10_closure.go`
- **Ponteiros:** `aula109_ponteiros.go`, `aula111_exercicio2_ponteiro_struct.go`
- **JSON:** `aula113_json_marshal_unmarshal.go`, `aula114_json_marshal_ordenacao.go`, `aula115_json_unmarshal_desordenado.go`
- **Concorrência:** `aula126_goroutines_waitgroup.go`, `aula129_mutex.go`, `aula130_atomic.go`, `aula146_canais.go`, `aula149_select_canais.go`
- **Erros:** `aula161_erros.go`, `aula164_recover.go`, `aula166_exercicio3_erroespecial.go`
- **Testes:** `aula179_testes_benchmarks_test.go`, `aula180_testes_tabela.go`, `exemplo_teste_simples_test.go`
- **Ferramentas:** `aula182_fmt.go`, `aula182_vet.go`, `aula182_golint.go`

---

## Dicas
- Os arquivos seguem o padrão `aulaXX_...` ou `exercicio_...` para facilitar a busca.
- Consulte o código de cada arquivo para exemplos comentados e explicações.
- Sinta-se à vontade para modificar, adicionar ou remover exercícios conforme seu progresso! 