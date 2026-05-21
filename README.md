# 📚 Estudos em Go - Jornada Completa de Aprendizado

<div align="center">

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-4.13.4-00ADD8?style=for-the-badge&logo=go&logoColor=white)

**Repositório completo de estudos e projetos em Go - Do básico ao avançado**

[🎯 Visão Geral](#-visão-geral) • [📁 Estrutura](#-estrutura) • [🚀 Projetos](#-projetos) • [⚡ Goroutines](#4--goroutines--concorrência-aplicada-profissional) • [📡 Channels](#5--channels--canais-aplicados) • [🔌 Middleware](#6--middleware--http-aplicado) • [🚦 Rate-Limiting](#7--rate-limiting--limite-de-taxa-http) • [📖 Conceitos](#-conceitos) • [🔧 Como Usar](#-como-usar)

</div>

---

## 🎯 Visão Geral

Este repositório contém uma **jornada completa de aprendizado em Go**, desde os conceitos fundamentais até o desenvolvimento de APIs RESTful completas. Além dos projetos em `Estudos-Realizados/`, há **trilhas profissionais** ([Goroutines/](Goroutines/), [Channels/](Channels/), [Middleware/](Middleware/), [Rate-Limiting/](Rate-Limiting/)) com documentação em Markdown e exemplos executáveis.

### 📊 Estatísticas dos Estudos
- **3 Projetos de API** desenvolvidos (`Estudos-Realizados/`)
- **4 Trilhas aplicadas** ([`Goroutines/`](Goroutines/), [`Channels/`](Channels/), [`Middleware/`](Middleware/), [`Rate-Limiting/`](Rate-Limiting/))
- **100+ Exercícios** práticos no curso base
- **29 capítulos** de documentação + **29 exemplos** executáveis nas trilhas
- **3 APIs RESTful** implementadas
- **Arquiteturas diferentes** exploradas

---

## 📁 Estrutura do Repositório

```
Estudo-GO/
├── Goroutines/                   # Concorrência aplicada (docs + examples)
│   ├── docs/                     # 8 capítulos
│   └── examples/                 # 01 … 05
├── Channels/                     # Canais (docs + examples)
│   ├── docs/                     # 7 capítulos
│   └── examples/                 # 01 … 07
├── Middleware/                   # HTTP middleware (docs + examples)
│   ├── docs/                     # 7 capítulos
│   ├── internal/httpmw/          # logging, recover, chain, request ID
│   └── examples/                 # 01 … 10
├── Rate-Limiting/                # Limite de taxa HTTP (docs + examples)
│   ├── docs/                     # 7 capítulos
│   ├── internal/ratelimit/       # middleware global e por IP
│   └── examples/                 # 01 … 07
└── Estudos-Realizados/
    ├── Curso_Aprenda_GO/         # Fundamentos e exercícios práticos
    ├── Curso-API-GO/             # API Todo com Chi Router
    └── API-Kelche/               # API Fitness com Echo Framework
```

### Mapa rápido — trilhas aplicadas

| Trilha | README | Docs | Exemplos | Foco |
|--------|--------|------|----------|------|
| [Goroutines/](Goroutines/) | [README](Goroutines/README.md) | [8 capítulos](Goroutines/docs/README.md) | `01`–`05` | `context`, pool, pipeline, errgroup, shutdown |
| [Channels/](Channels/) | [README](Channels/README.md) | [7 capítulos](Channels/docs/README.md) | `01`–`07` | buffer, `close`, `select`, fan-in, armadilhas |
| [Middleware/](Middleware/) | [README](Middleware/README.md) | [7 capítulos](Middleware/docs/README.md) | `01`–`10` | `http.Handler`, chain, Chi/Echo bridge |
| [Rate-Limiting/](Rate-Limiting/) | [README](Rate-Limiting/README.md) | [7 capítulos](Rate-Limiting/docs/README.md) | `01`–`07` | token bucket, 429, per-IP, Chi/Echo |

**Ordem sugerida (concorrência + HTTP):** curso → Channels `01`–`05` → Goroutines `01`–`05` → Middleware `01`–`08` → Rate-Limiting `01`–`05` → Channels `06`–`07` + Middleware `09`–`10` + Rate-Limiting `06`–`07`.

---

## 🚀 Projetos e trilhas

### 1. 📚 **Curso_Aprenda_GO** - Fundamentos Completos
**Objetivo:** Aprender os conceitos fundamentais da linguagem Go através de exercícios práticos.

#### 🎯 **Conceitos Abordados:**

##### **🔤 Conceitos Básicos**
- **Hello World** - Primeiro contato com Go
- **Variáveis e Tipos** - Declaração e tipos de dados
- **Valores Zero** - Valores padrão dos tipos
- **Constantes e Iota** - Constantes e enumerações
- **Operadores** - Aritméticos, lógicos e de comparação
- **Controle de Fluxo** - Loops, condicionais, switch

##### **📊 Estruturas de Dados**
- **Arrays** - Coleções de tamanho fixo
- **Slices** - Coleções dinâmicas (fatiamento, make, append)
- **Maps** - Estruturas chave-valor
- **Structs** - Estruturas personalizadas
- **Interfaces** - Contratos para tipos

##### **⚡ Concorrência**
- **Goroutines** - Execução concorrente
- **Canais** - Comunicação entre goroutines ([trilha Channels](Channels/))
- **WaitGroups** - Sincronização
- **Mutex e Atomic** - Controle de acesso
- **Select** - Multiplexação de canais
- **Context** - Cancelamento e timeouts  
- **Aprofundamento:** trilhas [`Goroutines/`](Goroutines/) e [`Channels/`](Channels/)

##### **🛠️ Ferramentas e Testes**
- **Testes Unitários** - Testes automatizados
- **Benchmarks** - Medição de performance
- **go fmt** - Formatação de código
- **go vet** - Análise estática
- **godoc** - Documentação

#### 📁 **Arquivos Principais:**
- `Estudos.go` - Arquivo principal com 4213 linhas de código
- `exercicios/` - 100+ exercícios organizados por tema
- `README.md` - Documentação detalhada dos exercícios

---

### 2. 🚀 **Curso-API-GO** - API Todo com Chi Router
**Objetivo:** Desenvolver uma API RESTful completa para gerenciamento de tarefas.

#### 🏗️ **Arquitetura:**
```
Curso-API-GO/
├── 📄 main.go              # Ponto de entrada
├── 📁 handlers/            # Controladores HTTP
├── 📁 models/              # Modelos de dados
├── 📁 db/                  # Configuração do banco
├── 📁 Configs/             # Configurações
└── 📄 config.toml          # Arquivo de configuração
```

#### 🛠️ **Tecnologias:**
- **Go 1.21+** - Linguagem principal
- **Chi Router** - Roteamento HTTP
- **PostgreSQL** - Banco de dados
- **Viper** - Gerenciamento de configurações
- **Docker** - Containerização

#### ✨ **Funcionalidades:**
- ✅ **CRUD Completo** - Criar, ler, atualizar, deletar tarefas
- ✅ **Validação de Dados** - Validação de entrada
- ✅ **Configuração Flexível** - Arquivo TOML
- ✅ **Logs Estruturados** - Sistema de logging
- ✅ **Health Check** - Monitoramento
- ✅ **Docker Support** - Containerização

#### 📚 **Endpoints:**
```
GET    /todos          # Listar todas as tarefas
GET    /todos/:id      # Buscar tarefa por ID
POST   /todos          # Criar nova tarefa
PUT    /todos/:id      # Atualizar tarefa
DELETE /todos/:id      # Deletar tarefa
GET    /health         # Health check
```

---

### 3. 🏋️ **API-Kelche** - API Fitness com Echo
**Objetivo:** Desenvolver uma API para gerenciamento de fitness com usuários e medições.

#### 🏗️ **Arquitetura:**
```
fitness-api/
├── 📄 main.go                    # Ponto de entrada
├── 📁 internal/
│   ├── 📁 handlers/              # Controladores
│   ├── 📁 models/                # Modelos
│   └── 📁 repositories/          # Acesso a dados
├── 📁 db/                        # Configuração DB
└── 📄 schema.sql                 # Schema do banco
```

#### 🛠️ **Tecnologias:**
- **Go 1.24.4** - Linguagem principal
- **Echo Framework** - Framework web
- **PostgreSQL** - Banco de dados
- **Repository Pattern** - Padrão de acesso a dados

#### ✨ **Funcionalidades:**
- 👥 **Gestão de Usuários** - CRUD de usuários
- 📊 **Medições de Fitness** - Peso, altura, gordura corporal
- 🔐 **Validação de Dados** - Validação de entrada
- 📈 **Relatórios** - Histórico de medições
- 🏗️ **Arquitetura Limpa** - Separação de responsabilidades

#### 📚 **Endpoints:**
```
GET    /users                    # Listar usuários
GET    /users/:id                # Buscar usuário
POST   /users                    # Criar usuário
PUT    /users/:id                # Atualizar usuário
DELETE /users/:id                # Deletar usuário

GET    /measurements             # Listar medições
GET    /measurements/:id         # Buscar medição
POST   /measurements             # Criar medição
PUT    /measurements/:id         # Atualizar medição
DELETE /measurements/:id         # Deletar medição
```

---

### 4. ⚡ **Goroutines** — Concorrência aplicada (profissional)

**Objetivo:** Ir além dos exercícios do curso (race, mutex, 50k goroutines) e fixar padrões usados em **APIs**, **workers** e **batch/ETL**: `context`, worker pool, pipeline, `errgroup`, shutdown gracioso.

#### 📁 **Estrutura:**
```
Goroutines/
├── README.md
├── go.mod
├── docs/                         # Trilha 01–08
└── examples/
    ├── 01-context-cancel/
    ├── 02-worker-pool/
    ├── 03-errgroup/
    ├── 04-pipeline/
    └── 05-graceful-shutdown/
```

#### 📚 **Documentação:** [Goroutines/docs/README.md](Goroutines/docs/README.md)

| # | Tema |
|---|------|
| 1 | Por que não spawn infinito |
| 2 | Context e cancelamento |
| 3 | WaitGroup e errgroup |
| 4 | Worker pool e backpressure |
| 5 | Pipeline e fan-out / fan-in |
| 6 | Mutex, atomic e canais |
| 7 | Leaks e debugging (`-race`, pprof) |
| 8 | Shutdown gracioso |

**Relação com o curso:** fundamentos em `Curso_Aprenda_GO` (aulas 126, 139, canais); esta trilha foca **uso em sistema**.

---

### 5. 📡 **Channels** — Canais aplicados (profissional)

**Objetivo:** Dominar **semântica de canais**: buffer, `close`, direção, `select`, fan-in e armadilhas — complemento direto de [Goroutines/](Goroutines/).

#### 📁 **Estrutura:**
```
Channels/
├── README.md
├── go.mod
├── docs/                         # Trilha 01–07
└── examples/
    ├── 01-basics/
    ├── 02-buffered/
    ├── 03-close-range/
    ├── 04-directional/
    ├── 05-select-timeout/
    ├── 06-select-quit/
    └── 07-fan-in/
```

#### 📚 **Documentação:** [Channels/docs/README.md](Channels/docs/README.md)

| # | Tema |
|---|------|
| 1 | O que é um canal |
| 2 | Buffered vs unbuffered |
| 3 | Close, range e sinais |
| 4 | Canais direcionais |
| 5 | Select |
| 6 | Padrões em produção |
| 7 | Armadilhas e leaks |

**Ordem sugerida:** Channels `01`–`05` → Goroutines pool/pipeline → Channels `06`–`07`.

---

### 6. 🔌 **Middleware** — HTTP aplicado (profissional)

**Objetivo:** Entender middleware como `func(http.Handler) http.Handler`, construir logging/recover/request ID do zero e reutilizar o mesmo código com **Chi** e **Echo**.

#### 📁 **Estrutura:**
```
Middleware/
├── README.md
├── go.mod
├── docs/                         # Trilha 01–07
├── internal/httpmw/              # Chain, Logging, Recover, RequestID
└── examples/
    ├── 01-hello-handler/ … 08-stdlib-server/
    ├── 09-chi-bridge/
    └── 10-echo-bridge/
```

#### 📚 **Documentação:** [Middleware/docs/README.md](Middleware/docs/README.md)

| # | Tema |
|---|------|
| 1 | Handler e ServeMux |
| 2 | Assinatura do middleware |
| 3 | Cadeia e ordem |
| 4 | ResponseWriter wrapper |
| 5 | Context e request ID |
| 6 | Chi, Echo e stdlib |
| 7 | Armadilhas |

**Relação com APIs:** [Curso-API-GO](Estudos-Realizados/Curso-API-GO/main.go) (Chi), [API-Kelche](Estudos-Realizados/API-Kelche/fitness-api/main.go) (Echo).

---

### 7. 🚦 **Rate-Limiting** — Limite de taxa HTTP (profissional)

**Objetivo:** Proteger APIs com limite de requisições por segundo, middleware **429**, limite global e por IP com `golang.org/x/time/rate`, e pontes para **Chi** e **Echo** (após [Middleware/](Middleware/)).

#### 📁 **Estrutura:**
```
Rate-Limiting/
├── README.md
├── go.mod
├── docs/                         # Trilha 01–07
├── internal/ratelimit/           # Middleware, PerIP, fixed window
└── examples/
    ├── 01-token-bucket-basico/
    ├── 02-middleware-global/
    ├── 03-middleware-per-ip/
    ├── 04-fixed-window-manual/
    ├── 05-stdlib-server/
    ├── 06-chi-bridge/
    └── 07-echo-bridge/
```

#### 📚 **Documentação:** [Rate-Limiting/docs/README.md](Rate-Limiting/docs/README.md)

| # | Tema |
|---|------|
| 1 | Por que rate limiting |
| 2 | Conceitos e algoritmos |
| 3 | Token bucket (`x/time/rate`) |
| 4 | Middleware e HTTP 429 |
| 5 | Limite por IP e chave |
| 6 | Chi, Echo e produção |
| 7 | Armadilhas (mapa em memória, cluster) |

**Relação com APIs:** encaixar na chain após recover/logging ([Middleware](Middleware/docs/03-cadeia-e-ordem.md)); [Curso-API-GO](Estudos-Realizados/Curso-API-GO/main.go) e [API-Kelche](Estudos-Realizados/API-Kelche/fitness-api/README.md) ainda sem rate limit — esta trilha mostra como adicionar.

---

## 📖 Conceitos Aprendidos

### 🔤 **Fundamentos da Linguagem**
1. **Sintaxe e Estrutura** - Como escrever código Go
2. **Tipos de Dados** - Int, float, string, bool, etc.
3. **Variáveis e Constantes** - Declaração e escopo
4. **Operadores** - Aritméticos, lógicos, de atribuição
5. **Controle de Fluxo** - if, else, switch, loops

### 📊 **Estruturas de Dados**
1. **Arrays e Slices** - Coleções de dados
2. **Maps** - Estruturas chave-valor
3. **Structs** - Estruturas personalizadas
4. **Interfaces** - Contratos para tipos
5. **Ponteiros** - Referências de memória

### ⚡ **Concorrência**
1. **Goroutines** - Execução concorrente (curso + [trilha Goroutines](Goroutines/))
2. **Canais** - Comunicação entre goroutines ([trilha Channels](Channels/))
3. **Sincronização** - WaitGroups, Mutex, Atomic
4. **Padrões** - Worker pools, pipeline, fan-out/fan-in, errgroup
5. **Context** - Cancelamento, timeout, shutdown gracioso
6. **Qualidade** - `go test -race`, prevenção de goroutine leak

### 🛠️ **Desenvolvimento Web**
1. **HTTP/HTTPS** - Protocolos web
2. **REST APIs** - Arquitetura RESTful
3. **JSON** - Serialização de dados
4. **Middleware** - Interceptadores HTTP ([trilha Middleware](Middleware/))
5. **Rate limiting** - Token bucket, 429, limite por IP ([trilha Rate-Limiting](Rate-Limiting/))
6. **Roteamento** - Gerenciamento de rotas

### 🗄️ **Banco de Dados**
1. **PostgreSQL** - Banco relacional
2. **SQL** - Linguagem de consulta
3. **Drivers** - Conexão com banco
4. **Migrations** - Controle de schema
5. **ORM Patterns** - Padrões de acesso a dados

### 🧪 **Testes e Qualidade**
1. **Testes Unitários** - Testes automatizados
2. **Benchmarks** - Medição de performance
3. **Testes em Tabela** - Testes parametrizados
4. **Cobertura** - Cobertura de testes
5. **Ferramentas** - go fmt, go vet, golint

---

## 🔧 Como Usar

### 📋 **Pré-requisitos**
- **Go 1.21+** instalado (trilhas `Goroutines/` e `Channels/` usam módulos próprios; `go mod tidy` em cada pasta)
- **PostgreSQL** (para os projetos de API em `Estudos-Realizados/`)
- **Git** para controle de versão
- **Editor de código** (VS Code, GoLand, etc.)

Todos os comandos abaixo assumem que você está na raiz **`Estudo-GO/`**.

### 🚀 **Executando os estudos**

#### **Curso_Aprenda_GO**
```bash
# Navegar para o diretório
cd Estudos-Realizados/Curso_Aprenda_GO

# Executar um exercício específico
go run exercicios/aula2_hello_world.go

# Executar todos os testes
go test ./exercicios

# Executar o arquivo principal
go run Estudos.go
```

#### **Curso-API-GO**
```bash
# Navegar para o diretório
cd Estudos-Realizados/Curso-API-GO

# Instalar dependências
go mod tidy

# Configurar banco de dados
cp config.example.toml config.toml
# Editar config.toml com suas configurações

# Executar a API
go run main.go
```

#### **API-Kelche**
```bash
# Navegar para o diretório
cd Estudos-Realizados/API-Kelche/fitness-api

# Instalar dependências
go mod tidy

# Configurar banco de dados
psql -d fitness_api -f schema.sql

# Executar a API
go run main.go
```

#### **Goroutines** (trilha de concorrência)
```bash
cd Goroutines
go mod tidy

# Exemplos (ordem sugerida)
go run ./examples/01-context-cancel
go run ./examples/02-worker-pool
go run ./examples/03-errgroup
go run ./examples/04-pipeline
go run ./examples/05-graceful-shutdown   # Ctrl+C para testar shutdown

go test -race ./...   # quando houver testes no módulo
```

Documentação: [Goroutines/README.md](Goroutines/README.md) · [docs/](Goroutines/docs/README.md)

#### **Channels** (canais)
```bash
cd Channels

go run ./examples/01-basics
go run ./examples/02-buffered
go run ./examples/03-close-range
go run ./examples/04-directional
go run ./examples/05-select-timeout
go run ./examples/06-select-quit
go run ./examples/07-fan-in

go run -race ./examples/07-fan-in   # opcional: detector de race
```

Documentação: [Channels/README.md](Channels/README.md) · [docs/](Channels/docs/README.md)

#### **Concorrência — leitura + código**
| Etapa | Onde |
|-------|------|
| 1 | [Channels/docs/01](Channels/docs/01-o-que-e-canal.md) … [05](Channels/docs/05-select.md) + exemplos `01`–`05` |
| 2 | [Goroutines/docs/02](Goroutines/docs/02-context-e-cancelamento.md) … [05](Goroutines/docs/05-pipeline-e-fan-out-fan-in.md) + exemplos `01`–`05` |
| 3 | [Channels/docs/06](Channels/docs/06-padroes-producao.md)–[07](Channels/docs/07-armadilhas-e-leaks.md) + `06`–`07` |
| 4 | [Goroutines/docs/07](Goroutines/docs/07-leaks-e-debugging.md)–[08](Goroutines/docs/08-shutdown-gracioso.md) |

#### **Middleware** (HTTP)
```bash
cd Middleware
go mod tidy
go build ./...

go run ./examples/08-stdlib-server
go run ./examples/09-chi-bridge
go run ./examples/10-echo-bridge
```

Documentação: [Middleware/README.md](Middleware/README.md) · [docs/](Middleware/docs/README.md)

#### **Rate-Limiting** (limite de taxa)
```bash
cd Rate-Limiting
go mod tidy
go build ./...

go run ./examples/01-token-bucket-basico
go run ./examples/05-stdlib-server

# Várias requisições até ver 429 (PowerShell):
1..20 | ForEach-Object { curl.exe -s -o NUL -w "%{http_code}\n" http://localhost:8080/ }

go run ./examples/06-chi-bridge
go run ./examples/07-echo-bridge
```

Documentação: [Rate-Limiting/README.md](Rate-Limiting/README.md) · [docs/](Rate-Limiting/docs/README.md)

---

## 📈 Progresso de Aprendizado

### 🎯 **Nível Básico** ✅
- [x] Sintaxe e estrutura da linguagem
- [x] Tipos de dados fundamentais
- [x] Controle de fluxo
- [x] Estruturas de dados básicas
- [x] Funções e métodos

### 🎯 **Nível Intermediário** ✅
- [x] Ponteiros e referências
- [x] Interfaces e polimorfismo
- [x] Tratamento de erros
- [x] JSON e serialização
- [x] Testes unitários

### 🎯 **Nível Avançado** ✅
- [x] Concorrência com goroutines (curso)
- [x] Canais e sincronização
- [x] Desenvolvimento de APIs REST
- [x] Integração com banco de dados
- [x] Arquitetura de software
- [x] Trilha **Goroutines** — context, pool, pipeline, errgroup, shutdown ([docs](Goroutines/docs/README.md))
- [x] Trilha **Channels** — buffer, close, select, fan-in ([docs](Channels/docs/README.md))
- [x] Trilha **Middleware** — chain, logging, recover, Chi/Echo ([docs](Middleware/docs/README.md))
- [x] Trilha **Rate-Limiting** — token bucket, 429, per-IP, Chi/Echo ([docs](Rate-Limiting/docs/README.md))

### 🎯 **Nível Expert** ✅
- [x] Padrões de projeto
- [x] Performance e otimização
- [x] Containerização
- [x] Monitoramento e logs
- [x] Boas práticas de desenvolvimento

---

## 🏆 Conquistas

### 📚 **Conhecimento Adquirido**
- ✅ Domínio completo da sintaxe Go
- ✅ Compreensão profunda de concorrência
- ✅ Desenvolvimento de APIs RESTful
- ✅ Integração com bancos de dados
- ✅ Testes e qualidade de código
- ✅ Arquitetura de software

### 🚀 **Projetos Concluídos**
- ✅ **100+ Exercícios** práticos (Curso Aprenda Go)
- ✅ **API Todo** com Chi Router
- ✅ **API Fitness** com Echo Framework
- ✅ **Trilha Goroutines** — 5 exemplos + 8 capítulos
- ✅ **Trilha Channels** — 7 exemplos + 7 capítulos
- ✅ **Trilha Middleware** — 10 exemplos + 7 capítulos
- ✅ **Trilha Rate-Limiting** — 7 exemplos + 7 capítulos
- ✅ **Padrões de projeto** implementados
- ✅ **Testes automatizados** escritos

### 🛠️ **Ferramentas Dominadas**
- ✅ **Go Modules** - Gerenciamento de dependências
- ✅ **Chi Router** - Roteamento HTTP
- ✅ **Echo Framework** - Framework web
- ✅ **PostgreSQL** - Banco de dados
- ✅ **Docker** - Containerização
- ✅ **Testes** - Testes unitários e benchmarks

---

## 📚 Recursos Adicionais

### 🔗 **Links Úteis**
- [Documentação Oficial do Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Web Examples](https://gowebexamples.com/)

### 📖 **Livros Recomendados**
- "The Go Programming Language" - Alan Donovan & Brian Kernighan
- "Go in Action" - William Kennedy
- "Concurrency in Go" - Katherine Cox-Buday
- Trilhas locais: [Goroutines/docs](Goroutines/docs/README.md) · [Channels/docs](Channels/docs/README.md) · [Middleware/docs](Middleware/docs/README.md) · [Rate-Limiting/docs](Rate-Limiting/docs/README.md)

### 🎥 **Cursos Online**
- [Aprenda Go](https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)
- [Go Programming Tutorials](https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX)

---

## 🤝 Contribuição

Este repositório é pessoal e documenta minha jornada de aprendizado em Go. Sinta-se à vontade para:

- 🔍 **Explorar** os códigos e exemplos
- 📝 **Sugerir** melhorias ou correções
- 🚀 **Inspirar-se** para seus próprios projetos
- 📚 **Usar** como referência de estudo

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

<div align="center">

**Desenvolvido com ❤️ durante a jornada de aprendizado em Go**

*"A melhor forma de aprender é fazendo"*

</div> 