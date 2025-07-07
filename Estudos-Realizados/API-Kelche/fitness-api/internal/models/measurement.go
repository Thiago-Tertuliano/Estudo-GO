package models // Pacote models contém as estruturas de dados (structs) da aplicação

/*
Este arquivo (measurement.go) define a estrutura de dados para medições de fitness.
As medições representam dados físicos dos usuários como peso, altura e gordura corporal.

Conceitos importantes:
- Relacionamento: cada medição pertence a um usuário (user_id)
- Tipos numéricos: float64 para valores decimais precisos
- Timestamps: para rastrear quando a medição foi feita
*/

import (
	"time" // Biblioteca para trabalhar com datas e horários
)

// Measurements representa uma medição de fitness de um usuário
// Esta struct armazena dados físicos como peso, altura e gordura corporal
type Measurements struct {
	// ID único da medição (chave primária no banco de dados)
	// int = número inteiro
	Id int `json:"id"`

	// ID do usuário que fez a medição (chave estrangeira)
	// Relaciona esta medição com um usuário específico
	// No banco de dados, user_id referencia a tabela users(id)
	UserId int `json:"user_id"`

	// Peso em quilogramas (kg)
	// float64 = número decimal com alta precisão
	// Exemplo: 75.5 kg
	Weight float64 `json:"weight"`

	// Altura em metros (m)
	// float64 = número decimal com alta precisão
	// Exemplo: 1.75 m (175 cm)
	Height float64 `json:"height"`

	// Percentual de gordura corporal (%)
	// float64 = número decimal com alta precisão
	// Exemplo: 15.2% de gordura corporal
	BodyFat float64 `json:"body_fat"`

	// Data e hora em que a medição foi registrada
	// time.Time = tipo específico para datas em Go
	// Útil para acompanhar a evolução do usuário ao longo do tempo
	CreatedAt time.Time `json:"created_at"`
}

/*
EXEMPLO DE JSON DE UMA MEDIÇÃO:

{
  "id": 1,
  "user_id": 5,
  "weight": 75.5,
  "height": 1.75,
  "body_fat": 15.2,
  "created_at": "2024-01-15T14:30:00Z"
}

RELACIONAMENTO COM USUÁRIOS:

- user_id = 5 significa que esta medição pertence ao usuário com ID 5
- Uma medição sempre pertence a um usuário (relacionamento obrigatório)
- Um usuário pode ter várias medições (relacionamento 1:N)

EXEMPLO DE USO:

measurement := Measurements{
    UserId: 1,
    Weight: 70.0,
    Height: 1.70,
    BodyFat: 12.5,
}

// Quando convertido para JSON:
// {"user_id":1,"weight":70,"height":1.7,"body_fat":12.5}

CÁLCULOS ÚTEIS COM ESTES DADOS:

1. IMC (Índice de Massa Corporal):
   IMC = peso / (altura * altura)
   Exemplo: 70 / (1.70 * 1.70) = 24.22

2. Peso ideal (fórmula de Lorentz):
   Peso ideal = (altura - 100) - ((altura - 150) / 4)
   Exemplo: (170 - 100) - ((170 - 150) / 4) = 70 - 5 = 65 kg

3. Massa magra estimada:
   Massa magra = peso * (1 - gordura_corporal / 100)
   Exemplo: 70 * (1 - 12.5 / 100) = 70 * 0.875 = 61.25 kg
*/