# 🚀 API Todo em Go

<div align="center">

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)

**Uma API RESTful completa e robusta para gerenciamento de tarefas desenvolvida em Go com PostgreSQL**

[📖 Documentação](#-documentação) • [🚀 Começando](#-começando) • [🔧 Configuração](#-configuração) • [📚 API Reference](#-api-reference) • [🧪 Testes](#-testes)

</div>

---

## 📋 Índice

- [🎯 Sobre o Projeto](#-sobre-o-projeto)
- [✨ Funcionalidades](#-funcionalidades)
- [🛠️ Tecnologias](#️-tecnologias)
- [📦 Pré-requisitos](#-pré-requisitos)
- [🚀 Começando](#-começando)
- [🔧 Configuração](#-configuração)
- [📚 API Reference](#-api-reference)
- [🧪 Testes](#-testes)
- [📁 Estrutura do Projeto](#-estrutura-do-projeto)
- [🐳 Docker](#-docker)
- [🔍 Monitoramento](#-monitoramento)
- [🐛 Solução de Problemas](#-solução-de-problemas)
- [🤝 Contribuição](#-contribuição)
- [📄 Licença](#-licença)

---

## 🎯 Sobre o Projeto

Este projeto implementa uma **API RESTful completa** para gerenciamento de tarefas (Todo List) utilizando **Go** como linguagem de programação e **PostgreSQL** como banco de dados. A aplicação segue as melhores práticas de desenvolvimento, incluindo:

- ✅ **Arquitetura limpa** com separação de responsabilidades
- ✅ **Configuração flexível** usando Viper
- ✅ **Roteamento eficiente** com Chi Router
- ✅ **Operações CRUD completas**
- ✅ **Validação de dados**
- ✅ **Tratamento de erros robusto**
- ✅ **Documentação completa**

### 🎯 Objetivos

- Demonstrar boas práticas de desenvolvimento em Go
- Implementar uma API RESTful completa e funcional
- Fornecer uma base sólida para projetos maiores
- Facilitar o aprendizado de Go e APIs

---

## ✨ Funcionalidades

### 🔧 Funcionalidades Principais

- **📝 Gerenciamento de Tarefas**: Criar, ler, atualizar e deletar tarefas
- **🔍 Busca Avançada**: Buscar tarefas por ID
- **📋 Listagem Completa**: Listar todas as tarefas com ordenação
- **✅ Status de Tarefas**: Marcar tarefas como concluídas ou pendentes
- **⏰ Timestamps**: Controle automático de criação e atualização
- **🔒 Validação**: Validação de dados de entrada
- **📊 Health Check**: Endpoint para verificar status da API

### 🚀 Funcionalidades Técnicas

- **🔄 RESTful API**: Seguindo padrões REST
- **📡 HTTP/HTTPS**: Suporte completo a protocolos HTTP
- **🔐 Configuração Segura**: Gerenciamento seguro de configurações
- **📝 Logs Estruturados**: Sistema de logging organizado
- **⚡ Performance**: Otimizado para alta performance
- **🔧 Middleware**: Middleware para logging, recuperação e timeout

---

## 🛠️ Tecnologias

### 🎯 Linguagem e Framework

| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| **Go** | 1.21+ | Linguagem de programação principal |
| **Chi Router** | v5.0.10 | Router HTTP leve e rápido |
| **Viper** | v1.17.0 | Gerenciamento de configurações |

### 🗄️ Banco de Dados

| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| **PostgreSQL** | 15+ | Banco de dados relacional |
| **lib/pq** | v1.10.9 | Driver PostgreSQL para Go |

### 🐳 Infraestrutura

| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| **Docker** | 20+ | Containerização |
| **Docker Compose** | 2+ | Orquestração de containers |

---

## 📦 Pré-requisitos

### 💻 Desenvolvimento

- **Go** 1.21 ou superior
- **Git** para controle de versão
- **Editor de código** (VS Code, GoLand, Vim, etc.)

### 🗄️ Banco de Dados

- **PostgreSQL** 15 ou superior
- **Ou Docker** para containerização

### 🔧 Ferramentas Opcionais

- **Postman** ou **Insomnia** para testar a API
- **Docker Desktop** para containerização
- **pgAdmin** para gerenciar PostgreSQL

---

## 🚀 Começando

### 1️⃣ Clone o Repositório

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/curso-api-go.git

# Entre no diretório
cd curso-api-go

# Verifique se está no diretório correto
ls -la
```

### 2️⃣ Instale as Dependências

```bash
# Inicialize o módulo Go (se necessário)
go mod init curso-api-go

# Baixe as dependências
go mod tidy

# Verifique as dependências
go mod graph
```

### 3️⃣ Configure o Ambiente

```bash
# Copie o arquivo de configuração de exemplo
cp config.example.toml config.toml

# Edite as configurações conforme necessário
nano config.toml
```

---

## 🔧 Configuração

### 🗄️ Configuração do Banco de Dados

#### Opção A: PostgreSQL Local

```bash
# Instale o PostgreSQL (Ubuntu/Debian)
sudo apt update
sudo apt install postgresql postgresql-contrib

# Inicie o serviço
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Acesse o PostgreSQL
sudo -u postgres psql
```

#### Opção B: Docker (Recomendado)

```bash
# Execute o container PostgreSQL
docker run --name postgres-todo \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=api_todo \
  -p 5432:5432 \
  -d postgres:15

# Verifique se está rodando
docker ps

# Acesse o container
docker exec -it postgres-todo psql -U postgres
```

### 🗃️ Criação do Banco de Dados

```sql
-- Conectar ao PostgreSQL
psql -U postgres

-- Criar banco de dados
CREATE DATABASE api_todo;

-- Conectar ao banco de dados
\c api_todo;

-- Criar usuário
CREATE USER user_todo WITH PASSWORD 'password123';

-- Conceder permissões
GRANT ALL PRIVILEGES ON DATABASE api_todo TO user_todo;

-- Criar tabela
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Conceder permissões na tabela
GRANT ALL PRIVILEGES ON TABLE todos TO user_todo;
GRANT USAGE, SELECT ON SEQUENCE todos_id_seq TO user_todo;

-- Verificar a tabela
\d todos;

-- Sair do PostgreSQL
\q
```

### ⚙️ Configuração da Aplicação

Edite o arquivo `config.toml`:

```toml
[api]
port = ":8080"

[database]
host = "localhost"        # Host do PostgreSQL
port = "5432"            # Porta do PostgreSQL
user = "user_todo"       # Usuário do banco
password = "password123" # Senha do banco
dbname = "api_todo"      # Nome do banco
sslmode = "disable"      # SSL mode
```

### 🚀 Executando a Aplicação

```bash
# Executar em modo desenvolvimento
go run main.go

# Ou compilar e executar
go build -o api-todo main.go
./api-todo

# Verificar se está rodando
curl http://localhost:8080/health
```

---

## 📚 API Reference

### 🔗 Base URL

```
http://localhost:8080
```

### 📋 Endpoints

| Método | Endpoint | Descrição | Status Code |
|--------|----------|-----------|-------------|
| `GET` | `/health` | Health check da API | `200` |
| `GET` | `/api/todos` | Listar todas as tarefas | `200` |
| `POST` | `/api/todos` | Criar nova tarefa | `201` |
| `GET` | `/api/todos/{id}` | Buscar tarefa por ID | `200` |
| `PUT` | `/api/todos/{id}` | Atualizar tarefa | `200` |
| `DELETE` | `/api/todos/{id}` | Deletar tarefa | `200` |

### 📝 Modelo de Dados

#### Todo Object

```json
{
  "id": 1,
  "title": "Estudar Go",
  "description": "Aprender Go e APIs REST",
  "done": false,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### 🔧 Exemplos de Uso

#### 1. Health Check

```bash
curl -X GET http://localhost:8080/health
```

**Resposta:**
```json
"API está funcionando!"
```

#### 2. Criar Tarefa

```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Estudar Go",
    "description": "Aprender Go e APIs REST",
    "done": false
  }'
```

**Resposta:**
```json
{
  "id": 1,
  "title": "Estudar Go",
  "description": "Aprender Go e APIs REST",
  "done": false,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

#### 3. Listar Todas as Tarefas

```bash
curl -X GET http://localhost:8080/api/todos
```

**Resposta:**
```json
[
  {
    "id": 1,
    "title": "Estudar Go",
    "description": "Aprender Go e APIs REST",
    "done": false,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  {
    "id": 2,
    "title": "Fazer exercícios",
    "description": "Treinar na academia",
    "done": true,
    "created_at": "2024-01-15T11:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
]
```

#### 4. Buscar Tarefa por ID

```bash
curl -X GET http://localhost:8080/api/todos/1
```

**Resposta:**
```json
{
  "id": 1,
  "title": "Estudar Go",
  "description": "Aprender Go e APIs REST",
  "done": false,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

#### 5. Atualizar Tarefa

```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Estudar Go - Atualizado",
    "description": "Aprender Go e APIs REST",
    "done": true
  }'
```

**Resposta:**
```json
{
  "message": "todo atualizado com sucesso"
}
```

#### 6. Deletar Tarefa

```bash
curl -X DELETE http://localhost:8080/api/todos/1
```

**Resposta:**
```json
{
  "message": "todo deletado com sucesso"
}
```

### ⚠️ Códigos de Erro

| Código | Descrição |
|--------|-----------|
| `400` | Bad Request - Dados inválidos |
| `404` | Not Found - Recurso não encontrado |
| `405` | Method Not Allowed - Método não permitido |
| `500` | Internal Server Error - Erro interno do servidor |

---

## 🧪 Testes

### 🚀 Testando com cURL

```bash
# Script de teste completo
#!/bin/bash

BASE_URL="http://localhost:8080"

echo "🧪 Iniciando testes da API..."

# Health check
echo "1. Testando health check..."
curl -s $BASE_URL/health
echo -e "\n"

# Criar tarefa
echo "2. Criando tarefa..."
CREATE_RESPONSE=$(curl -s -X POST $BASE_URL/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Teste API", "description": "Testando a API", "done": false}')
echo $CREATE_RESPONSE
echo -e "\n"

# Extrair ID da resposta
ID=$(echo $CREATE_RESPONSE | grep -o '"id":[0-9]*' | cut -d':' -f2)

# Listar todas
echo "3. Listando todas as tarefas..."
curl -s $BASE_URL/api/todos
echo -e "\n"

# Buscar por ID
echo "4. Buscando tarefa por ID ($ID)..."
curl -s $BASE_URL/api/todos/$ID
echo -e "\n"

# Atualizar
echo "5. Atualizando tarefa..."
curl -s -X PUT $BASE_URL/api/todos/$ID \
  -H "Content-Type: application/json" \
  -d '{"title": "Teste API - Atualizado", "description": "Testando a API", "done": true}'
echo -e "\n"

# Deletar
echo "6. Deletando tarefa..."
curl -s -X DELETE $BASE_URL/api/todos/$ID
echo -e "\n"

echo "✅ Testes concluídos!"
```

### 🧪 Testando com Postman

1. **Importe a coleção** (se disponível)
2. **Configure as variáveis de ambiente**
3. **Execute os testes** em sequência

### 🧪 Testando com Insomnia

1. **Crie um novo projeto**
2. **Adicione as requisições** conforme os exemplos
3. **Configure o ambiente** com a base URL
4. **Execute os testes**

---

## 📁 Estrutura do Projeto

```
curso-api-go/
├── 📁 configs/
│   └── 📄 config.go              # Configurações da aplicação
├── 📁 DB/
│   └── 📄 connection.go          # Conexão com banco de dados
├── 📁 models/
│   ├── 📄 entities.go            # Entidades do banco
│   ├── 📄 insert.go              # Operação Create
│   ├── 📄 get.go                 # Operação Read (um registro)
│   ├── 📄 getAll.go              # Operação Read (todos)
│   ├── 📄 update.go              # Operação Update
│   └── 📄 delete.go              # Operação Delete
├── 📁 handlers/
│   ├── 📄 create.go              # Handler para criar
│   ├── 📄 get.go                 # Handler para buscar
│   ├── 📄 list.go                # Handler para listar
│   ├── 📄 update.go              # Handler para atualizar
│   └── 📄 delete.go              # Handler para deletar
├── 📄 main.go                    # Arquivo principal
├── 📄 config.toml                # Configurações
├── 📄 config.example.toml        # Exemplo de configuração
├── 📄 go.mod                     # Dependências Go
├── 📄 go.sum                     # Checksums das dependências
├── 📄 .gitignore                 # Arquivos ignorados pelo Git
├── 📄 README.md                  # Este arquivo
└── 📄 LICENSE                    # Licença do projeto
```

### 📋 Descrição dos Pacotes

| Pacote | Descrição |
|--------|-----------|
| `configs` | Gerenciamento de configurações da aplicação |
| `DB` | Conexão e configuração do banco de dados |
| `models` | Entidades e operações CRUD do banco |
| `handlers` | Manipuladores das requisições HTTP |

---

## 🐳 Docker

### 🚀 Executando com Docker

```bash
# Build da imagem
docker build -t curso-api-go .

# Executar container
docker run -p 8080:8080 --name api-todo curso-api-go

# Executar com variáveis de ambiente
docker run -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_USER=user_todo \
  -e DB_PASSWORD=password123 \
  -e DB_NAME=api_todo \
  --name api-todo curso-api-go
```

### 🐳 Docker Compose

```yaml
# docker-compose.yml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - postgres
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=user_todo
      - DB_PASSWORD=password123
      - DB_NAME=api_todo

  postgres:
    image: postgres:15
    environment:
      - POSTGRES_DB=api_todo
      - POSTGRES_USER=user_todo
      - POSTGRES_PASSWORD=password123
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

```bash
# Executar com Docker Compose
docker-compose up -d

# Verificar logs
docker-compose logs -f

# Parar serviços
docker-compose down
```

---

## 🔍 Monitoramento

### 📊 Métricas de Performance

A API inclui middleware para monitoramento:

- **Request ID**: Rastreamento único de requisições
- **Logging**: Logs estruturados de todas as requisições
- **Timeout**: Timeout configurável para requisições
- **Recovery**: Recuperação automática de panics

### 📈 Health Check

```bash
# Verificar status da API
curl http://localhost:8080/health

# Verificar com mais detalhes
curl -v http://localhost:8080/health
```

### 📝 Logs

Os logs incluem:
- **Timestamp** da requisição
- **Request ID** único
- **Método HTTP** e **endpoint**
- **Status code** da resposta
- **Tempo de resposta**
- **IP do cliente**

---

## 🐛 Solução de Problemas

### ❌ Problemas Comuns

#### 1. Erro de Conexão com Banco

```bash
# Verificar se PostgreSQL está rodando
sudo systemctl status postgresql

# Verificar conectividade
telnet localhost 5432

# Testar conexão manual
psql -h localhost -U user_todo -d api_todo
```

#### 2. Erro de Dependências

```bash
# Limpar cache do Go
go clean -modcache

# Baixar dependências novamente
go mod tidy

# Verificar versão do Go
go version
```

#### 3. Porta Já em Uso

```bash
# Verificar processos na porta
lsof -i :8080

# Matar processo
kill -9 <PID>

# Ou alterar porta no config.toml
```

#### 4. Erro de Permissões no Banco

```sql
-- Conectar como superusuário
sudo -u postgres psql

-- Verificar usuário
\du

-- Recriar usuário se necessário
DROP USER IF EXISTS user_todo;
CREATE USER user_todo WITH PASSWORD 'password123';
GRANT ALL PRIVILEGES ON DATABASE api_todo TO user_todo;
```

### 🔧 Debug

```bash
# Executar com debug
DEBUG=true go run main.go

# Verificar variáveis de ambiente
env | grep DB_

# Verificar configuração
cat config.toml
```

---

## 🤝 Contribuição

### 📋 Como Contribuir

1. **Fork o projeto**
2. **Crie uma branch** para sua feature
   ```bash
   git checkout -b feature/AmazingFeature
   ```
3. **Commit suas mudanças**
   ```bash
   git commit -m 'Add some AmazingFeature'
   ```
4. **Push para a branch**
   ```bash
   git push origin feature/AmazingFeature
   ```
5. **Abra um Pull Request**

### 📝 Padrões de Commit

Siga o padrão [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: adiciona nova funcionalidade
fix: corrige bug
docs: atualiza documentação
style: formatação de código
refactor: refatoração de código
test: adiciona ou corrige testes
chore: tarefas de manutenção
```

### 🧪 Testes

Antes de submeter um PR:

```bash
# Executar testes
go test ./...

# Verificar formatação
go fmt ./...

# Verificar linting
golangci-lint run

# Verificar cobertura
go test -cover ./...
```

### 📋 Checklist

- [ ] Código segue os padrões do projeto
- [ ] Testes foram adicionados/atualizados
- [ ] Documentação foi atualizada
- [ ] Não há conflitos de merge
- [ ] Build passa sem erros

---

## 📄 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

```
MIT License

Copyright (c) 2025 Thiago Matos Tertuliano

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

## 📞 Suporte

### 🆘 Como Obter Ajuda

- **📖 Documentação**: Leia este README completamente
- **🐛 Issues**: Abra uma [issue](https://github.com/thiago-tertuliano/curso-api-go/issues) no GitHub
- **💬 Discussões**: Use as [discussions](https://github.com/thiago-tertuliano/curso-api-go/discussions) do GitHub
- **📧 Email**: Entre em contato via email (thiagomatostertuliano@gmail.com)

### 🔗 Links Úteis

- [📚 Documentação do Go](https://golang.org/doc/)
- [🗄️ Documentação do PostgreSQL](https://www.postgresql.org/docs/)
- [🔄 Documentação do Chi Router](https://github.com/go-chi/chi)
- [⚙️ Documentação do Viper](https://github.com/spf13/viper)

---

<div align="center">

**⭐ Se este projeto te ajudou, considere dar uma estrela no GitHub!**

Feito com ❤️ por Thiago Matos Tertuliano (https://github.com/thiago-tertuliano)

</div> 
