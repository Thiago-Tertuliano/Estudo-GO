-- =====================================================
-- SCRIPT DE CRIAÇÃO DO BANCO DE DADOS - API DE FITNESS
-- =====================================================
-- 
-- Este arquivo contém todas as instruções SQL necessárias
-- para criar as tabelas e estruturas do banco de dados
-- da API de Fitness.
--
-- CONCEITOS IMPORTANTES:
-- - DDL (Data Definition Language): comandos para criar estruturas
-- - DML (Data Manipulation Language): comandos para manipular dados
-- - Constraints: regras que garantem integridade dos dados
-- - Índices: estruturas que melhoram a performance das consultas
-- - Relacionamentos: conexões entre tabelas (chaves estrangeiras)
--
-- EXECUÇÃO:
-- psql -d nome_do_banco -f schema.sql
-- =====================================================

-- =====================================================
-- TABELA: users (Usuários)
-- =====================================================
-- Esta tabela armazena informações dos usuários do sistema
-- Cada usuário pode ter múltiplas medições (relacionamento 1:N)
CREATE TABLE IF NOT EXISTS users (
    -- Chave primária: identifica unicamente cada usuário
    -- SERIAL = auto-incremento (1, 2, 3, ...)
    id SERIAL PRIMARY KEY,
    
    -- Nome completo do usuário
    -- VARCHAR(255) = texto com máximo 255 caracteres
    -- NOT NULL = campo obrigatório (não pode ser vazio)
    name VARCHAR(255) NOT NULL,
    
    -- Email do usuário (deve ser único)
    -- UNIQUE = não pode haver emails duplicados
    -- NOT NULL = campo obrigatório
    email VARCHAR(255) UNIQUE NOT NULL,
    
    -- Senha do usuário (em produção deve ser criptografada!)
    -- VARCHAR(255) = espaço suficiente para hash da senha
    -- NOT NULL = campo obrigatório
    password VARCHAR(255) NOT NULL,
    
    -- Data e hora de criação do registro
    -- TIMESTAMP = data e hora completa
    -- DEFAULT CURRENT_TIMESTAMP = valor padrão é o momento atual
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Data e hora da última atualização
    -- Útil para saber quando o usuário foi modificado
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================================================
-- TABELA: measurements (Medições)
-- =====================================================
-- Esta tabela armazena as medições de fitness dos usuários
-- Cada medição pertence a um usuário específico (relacionamento N:1)
CREATE TABLE IF NOT EXISTS measurements (
    -- Chave primária: identifica unicamente cada medição
    id SERIAL PRIMARY KEY,
    
    -- Chave estrangeira: referencia o usuário que fez a medição
    -- INTEGER = número inteiro
    -- NOT NULL = campo obrigatório
    -- REFERENCES users(id) = referencia a coluna id da tabela users
    -- ON DELETE CASCADE = se o usuário for deletado, suas medições também serão
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Peso em quilogramas (kg)
    -- DECIMAL(5,2) = número decimal com 5 dígitos total, 2 após a vírgula
    -- Exemplo: 75.50 kg
    weight DECIMAL(5,2),
    
    -- Altura em metros (m)
    -- DECIMAL(5,2) = número decimal com 5 dígitos total, 2 após a vírgula
    -- Exemplo: 1.75 m
    height DECIMAL(5,2),
    
    -- Percentual de gordura corporal (%)
    -- DECIMAL(5,2) = número decimal com 5 dígitos total, 2 após a vírgula
    -- Exemplo: 15.20%
    body_fat DECIMAL(5,2),
    
    -- Data e hora em que a medição foi registrada
    -- TIMESTAMP = data e hora completa
    -- DEFAULT CURRENT_TIMESTAMP = valor padrão é o momento atual
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================================================
-- ÍNDICES PARA MELHORAR PERFORMANCE
-- =====================================================
-- Índices são estruturas que aceleram consultas no banco
-- São criados automaticamente para chaves primárias
-- Índices adicionais melhoram a performance de consultas específicas

-- Índice para busca rápida de usuários por email
-- Útil para login e verificação de emails únicos
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Índice para busca rápida de medições por usuário
-- Útil para listar todas as medições de um usuário específico
CREATE INDEX IF NOT EXISTS idx_measurements_user_id ON measurements(user_id);

-- Índice para busca rápida de medições por data
-- Útil para consultas por período de tempo
CREATE INDEX IF NOT EXISTS idx_measurements_created_at ON measurements(created_at);

-- =====================================================
-- COMENTÁRIOS EXPLICATIVOS
-- =====================================================

/*
CONCEITOS DE BANCO DE DADOS RELACIONAL:

1. TABELAS:
   - São como "planilhas" que organizam dados
   - Cada linha é um registro
   - Cada coluna é um campo/atributo

2. CHAVES PRIMÁRIAS (PRIMARY KEY):
   - Identificam unicamente cada registro
   - Não podem ser duplicadas
   - Não podem ser nulas

3. CHAVES ESTRANGEIRAS (FOREIGN KEY):
   - Referenciam chaves primárias de outras tabelas
   - Criam relacionamentos entre tabelas
   - Garantem integridade referencial

4. RELACIONAMENTOS:
   - 1:1 = Um para um
   - 1:N = Um para muitos (um usuário tem várias medições)
   - N:N = Muitos para muitos (precisa de tabela intermediária)

5. CONSTRAINTS (RESTRIÇÕES):
   - NOT NULL = campo obrigatório
   - UNIQUE = valor único
   - DEFAULT = valor padrão
   - CHECK = validação personalizada

6. ÍNDICES:
   - Aceleram consultas
   - Consomem espaço em disco
   - São criados automaticamente para chaves primárias

EXEMPLOS DE CONSULTAS ÚTEIS:

1. Buscar usuário com suas medições:
   SELECT u.name, m.weight, m.created_at 
   FROM users u 
   JOIN measurements m ON u.id = m.user_id 
   WHERE u.id = 1;

2. Calcular IMC de um usuário:
   SELECT u.name, m.weight, m.height, 
          m.weight / (m.height * m.height) as imc
   FROM users u 
   JOIN measurements m ON u.id = m.user_id 
   WHERE u.id = 1;

3. Última medição de cada usuário:
   SELECT u.name, m.weight, m.created_at
   FROM users u
   JOIN measurements m ON u.id = m.user_id
   WHERE m.created_at = (
       SELECT MAX(created_at) 
       FROM measurements 
       WHERE user_id = u.id
   );

4. Evolução do peso de um usuário:
   SELECT weight, created_at
   FROM measurements
   WHERE user_id = 1
   ORDER BY created_at ASC;

VALIDAÇÕES SUGERIDAS (TRIGGERS):

1. Verificar se peso está em range realista (30-300 kg)
2. Verificar se altura está em range realista (1.0-2.5 m)
3. Verificar se gordura corporal está em range realista (5-50%)
4. Evitar medições duplicadas no mesmo dia

MELHORIAS FUTURAS:

1. Adicionar campo updated_at na tabela measurements
2. Implementar soft delete (campo deleted_at)
3. Adicionar mais campos: circunferência, massa muscular, etc.
4. Criar tabela de metas/goals
5. Implementar sistema de logs de auditoria
*/

-- =====================================================
-- MENSAGEM DE SUCESSO
-- =====================================================
-- Este comentário será exibido quando o script for executado com sucesso
SELECT '✅ Banco de dados criado com sucesso!' as status; 