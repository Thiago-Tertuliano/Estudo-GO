# 🏋️ API de Fitness - Go + Echo + PostgreSQL

## 📚 Visão Geral Educativa

Esta é uma **API REST completa** desenvolvida em **Go** que demonstra os conceitos fundamentais de desenvolvimento web backend. A aplicação gerencia usuários e suas medições de fitness (peso, altura, gordura corporal), implementando um sistema completo de **CRUD** (Create, Read, Update, Delete).

### 🎯 Objetivos de Aprendizado

Este projeto foi criado para ensinar:
- **Arquitetura de APIs REST** com Go
- **Padrões de projeto** (Repository Pattern, MVC)
- **Banco de dados relacional** com PostgreSQL
- **Framework web** (Echo)
- **Boas práticas** de desenvolvimento
- **Estruturação de código** em Go

---

## 🏗️ Arquitetura da Aplicação

### 📁 Estrutura de Pacotes (Package Structure)

```
fitness-api/
├── 📄 main.go              # 🚀 Ponto de entrada da aplicação
├── 📄 database.go          # 🗄️ Configuração do banco de dados
├── 📄 user.go              # 👤 Modelo de dados do usuário
├── 📄 measurement.go       # 📊 Modelo de dados das medições
├── 📄 handleUsers.go       # 🎮 Handlers para operações de usuários
├── 📄 handleMeasurements.go # 🎮 Handlers para operações de medições
├── 📄 usersDb.go           # 💾 Operações de banco para usuários
├── 📄 measuremenstDb.go    # 💾 Operações de banco para medições
├── 📄 rootHandlers.go      # 🏠 Handler da página inicial
├── 📄 schema.sql           # 🗃️ Script de criação das tabelas
├── 📄 go.mod               # 📦 Gerenciamento de dependências
└── 📄 README.md            # 📖 Este arquivo
```

### 🔄 Fluxo de Dados (Data Flow)

```
Cliente HTTP → Echo Router → Handler → Repository → PostgreSQL
     ↑                                                      ↓
     ← JSON Response ← Handler ← Repository ← PostgreSQL ←
```

### 🎨 Padrões Utilizados

1. **Repository Pattern**: Separa a lógica de acesso a dados
2. **MVC (Model-View-Controller)**: Organização do código
3. **Dependency Injection**: Injeção de dependências
4. **Error Handling**: Tratamento adequado de erros

---

## 🛠️ Tecnologias e Conceitos

### 🔧 Stack Tecnológico

| Tecnologia | Versão | Propósito |
|------------|--------|-----------|
| **Go** | 1.24.4 | Linguagem de programação |
| **Echo** | v4.13.4 | Framework web |
| **PostgreSQL** | - | Banco de dados relacional |
| **lib/pq** | v1.10.9 | Driver PostgreSQL para Go |

### 📚 Conceitos Fundamentais

#### 1. **API REST (Representational State Transfer)**
- **GET**: Buscar dados
- **POST**: Criar novos recursos
- **PUT**: Atualizar recursos existentes
- **DELETE**: Remover recursos

#### 2. **HTTP Status Codes**
- **200 OK**: Sucesso
- **201 Created**: Recurso criado
- **400 Bad Request**: Dados inválidos
- **404 Not Found**: Recurso não encontrado
- **500 Internal Server Error**: Erro do servidor

#### 3. **JSON (JavaScript Object Notation)**
Formato de dados usado para comunicação cliente-servidor:
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@email.com",
  "created_at": "2024-01-15T10:30:00Z"
}
```

#### 4. **SQL (Structured Query Language)**
Linguagem para manipular bancos de dados:
```sql
SELECT * FROM users WHERE id = 1;
INSERT INTO users (name, email) VALUES ('João', 'joao@email.com');
```

---

## 🚀 Instalação e Configuração

### 📋 Pré-requisitos

1. **Go 1.24.4+** instalado
2. **PostgreSQL** instalado e rodando
3. **Git** para clonar o repositório

### 🔧 Passo a Passo

#### 1. **Clone o Repositório**
```bash
git clone <url-do-repositorio>
cd fitness-api
```

#### 2. **Instale as Dependências**
```bash
go mod tidy
```

#### 3. **Configure o Banco de Dados**

**Crie um banco PostgreSQL:**
```sql
CREATE DATABASE fitness_api;
```

**Execute o script de criação das tabelas:**
```bash
psql -d fitness_api -f schema.sql
```

#### 4. **Configure as Variáveis de Ambiente (Opcional)**

**Windows (CMD):**
```cmd
set DB_HOST=localhost
set DB_PORT=5432
set DB_USER=postgres
set DB_PASSWORD=sua_senha
set DB_NAME=fitness_api
```

**Linux/Mac:**
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=sua_senha
export DB_NAME=fitness_api
```

#### 5. **Execute a Aplicação**
```bash
go run main.go
```

A API estará disponível em: **http://localhost:8080**

---

## 📖 Documentação da API

### 🏠 Rota Raiz
```
GET /
```
**Resposta:** Mensagem de boas-vindas da API

### 👥 Endpoints de Usuários

#### Listar Todos os Usuários
```
GET /users
```
**Resposta:**
```json
[
  {
    "id": 1,
    "name": "João Silva",
    "email": "joao@email.com",
    "password": "senha123",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
]
```

#### Buscar Usuário por ID
```
GET /users/:id
```
**Exemplo:** `GET /users/1`

#### Criar Novo Usuário
```
POST /users
Content-Type: application/json

{
  "name": "Maria Santos",
  "email": "maria@email.com",
  "password": "senha456"
}
```

#### Atualizar Usuário
```
PUT /users/:id
Content-Type: application/json

{
  "name": "João Silva Atualizado",
  "email": "joao.novo@email.com",
  "password": "nova_senha"
}
```

#### Deletar Usuário
```
DELETE /users/:id
```

### 📊 Endpoints de Medições

#### Listar Todas as Medições
```
GET /measurements
```

#### Buscar Medição por ID
```
GET /measurements/:id
```

#### Criar Nova Medição
```
POST /measurements
Content-Type: application/json

{
  "user_id": 1,
  "weight": 75.5,
  "height": 1.75,
  "body_fat": 15.2
}
```

#### Atualizar Medição
```
PUT /measurements/:id
Content-Type: application/json

{
  "user_id": 1,
  "weight": 74.0,
  "height": 1.75,
  "body_fat": 14.8
}
```

#### Deletar Medição
```
DELETE /measurements/:id
```

---

## 🧪 Testando a API

### 🛠️ Ferramentas Recomendadas

1. **Postman**: Interface gráfica para testar APIs
2. **cURL**: Linha de comando
3. **Insomnia**: Alternativa ao Postman
4. **Navegador**: Para endpoints GET

### 📝 Exemplos de Teste

#### Via cURL

**Criar usuário:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@email.com",
    "password": "senha123"
  }'
```

**Listar usuários:**
```bash
curl http://localhost:8080/users
```

**Criar medição:**
```bash
curl -X POST http://localhost:8080/measurements \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "weight": 75.5,
    "height": 1.75,
    "body_fat": 15.2
  }'
```

#### Via Postman

1. **Crie uma nova coleção**
2. **Configure o base URL:** `http://localhost:8080`
3. **Adicione as requisições** conforme os endpoints acima
4. **Configure headers:** `Content-Type: application/json`

---

## 🗃️ Estrutura do Banco de Dados

### 📊 Diagrama das Tabelas

```
┌─────────────┐    ┌─────────────────┐
│    users    │    │   measurements  │
├─────────────┤    ├─────────────────┤
│ id (PK)     │◄───│ user_id (FK)    │
│ name        │    │ id (PK)         │
│ email       │    │ weight          │
│ password    │    │ height          │
│ created_at  │    │ body_fat        │
│ updated_at  │    │ created_at      │
└─────────────┘    └─────────────────┘
```

### 🔗 Relacionamentos

- **1:N**: Um usuário pode ter várias medições
- **FK**: `measurements.user_id` referencia `users.id`
- **CASCADE**: Se um usuário for deletado, suas medições também serão

### 📋 Script de Criação

```sql
-- Tabela de usuários
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de medições
CREATE TABLE IF NOT EXISTS measurements (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    weight DECIMAL(5,2),
    height DECIMAL(5,2),
    body_fat DECIMAL(5,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_measurements_user_id ON measurements(user_id);
CREATE INDEX IF NOT EXISTS idx_measurements_created_at ON measurements(created_at);
```

---

## 🧠 Conceitos de Programação em Go

### 📦 Pacotes (Packages)

```go
package main        // Pacote principal (executável)
package handlers    // Pacote para handlers HTTP
package models      // Pacote para estruturas de dados
package repositories // Pacote para acesso ao banco
package storage     // Pacote para configuração do banco
```

### 🏗️ Structs (Estruturas)

```go
type User struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### 🔄 Interfaces

```go
type UserRepository interface {
    Create(user User) (User, error)
    GetById(id int) (User, error)
    GetAll() ([]User, error)
    Update(user User, id int) (User, error)
    Delete(id int) error
}
```

### ⚡ Goroutines e Channels

```go
// Exemplo de goroutine para processamento assíncrono
go func() {
    // Processamento em background
}()

// Exemplo de channel para comunicação entre goroutines
ch := make(chan string)
go func() {
    ch <- "resultado"
}()
result := <-ch
```

---

## 🔒 Segurança e Boas Práticas

### 🛡️ Medidas de Segurança

1. **Prepared Statements**: Evitam SQL injection
2. **Validação de Entrada**: Sempre validar dados
3. **Criptografia de Senhas**: Usar bcrypt ou similar
4. **HTTPS**: Em produção, sempre usar HTTPS
5. **Rate Limiting**: Limitar requisições por IP

### 📋 Validações Implementadas

- ✅ Verificação de tipos de dados
- ✅ Validação de IDs numéricos
- ✅ Tratamento de erros de banco
- ✅ Códigos de status HTTP apropriados

### 🚨 Validações Pendentes

- ❌ Criptografia de senhas
- ❌ Validação de formato de email
- ❌ Autenticação e autorização
- ❌ Rate limiting
- ❌ Logs de auditoria

---

## 🚀 Melhorias e Próximos Passos

### 🔄 Funcionalidades Sugeridas

1. **Autenticação JWT**
   - Login/logout de usuários
   - Tokens de acesso
   - Middleware de autenticação

2. **Validação Avançada**
   - Validação de formato de email
   - Ranges realistas para medições
   - Validação de senhas fortes

3. **Funcionalidades de Fitness**
   - Cálculo de IMC
   - Gráficos de evolução
   - Metas de peso
   - Relatórios

4. **Performance**
   - Cache Redis
   - Paginação
   - Índices otimizados
   - Compressão de respostas

5. **Monitoramento**
   - Logs estruturados
   - Métricas de performance
   - Health checks
   - Alertas

### 🧪 Testes

1. **Testes Unitários**
   ```bash
   go test ./...
   ```

2. **Testes de Integração**
   - Testes com banco de dados real
   - Testes de endpoints HTTP

3. **Testes de Performance**
   - Benchmarks
   - Testes de carga

### 📚 Documentação

1. **Swagger/OpenAPI**
   - Documentação interativa
   - Testes direto na documentação

2. **Exemplos de Uso**
   - Casos de uso reais
   - Integração com frontend

---

## 🤝 Contribuindo

### 📝 Como Contribuir

1. **Fork** o repositório
2. **Crie** uma branch para sua feature
3. **Implemente** suas mudanças
4. **Teste** suas alterações
5. **Commit** suas mudanças
6. **Push** para a branch
7. **Abra** um Pull Request

### 🎯 Padrões de Código

- Use **gofmt** para formatação
- Siga as **convenções de nomenclatura** do Go
- Adicione **comentários** explicativos
- Escreva **testes** para novas funcionalidades

---

## 📞 Suporte e Comunidade

### 🆘 Como Obter Ajuda

1. **Issues do GitHub**: Reporte bugs e solicite features
2. **Documentação**: Leia este README e os comentários no código
3. **Comunidade Go**: Stack Overflow, Reddit r/golang
4. **Discord/Slack**: Comunidades de desenvolvedores Go

### 📚 Recursos de Aprendizado

- **[Go by Example](https://gobyexample.com/)**: Exemplos práticos
- **[Go Tour](https://tour.golang.org/)**: Tutorial interativo
- **[Echo Documentation](https://echo.labstack.com/)**: Documentação do Echo
- **[PostgreSQL Tutorial](https://www.postgresqltutorial.com/)**: Aprenda PostgreSQL

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

---

## 🙏 Agradecimentos

- **Go Team**: Pela linguagem incrível
- **Echo Framework**: Pelo framework web elegante
- **PostgreSQL**: Pelo banco de dados robusto
- **Comunidade Go**: Pelo suporte e recursos

---

**⭐ Se este projeto te ajudou, considere dar uma estrela no repositório!**

**🚀 Happy Coding!** 