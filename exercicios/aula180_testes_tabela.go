package calculo

import "testing"

func TestSomaTabela(t *testing.T) {
	testes := []struct {
		nome     string
		a, b     int
		esperado int
	}{
		{"soma positiva", 2, 3, 5},
		{"soma com zero", 0, 5, 5},
		{"soma negativa", -2, -3, -5},
		{"mistura", -1, 4, 3},
	}
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := Soma(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("esperado %d, obteve %d", tt.esperado, resultado)
			}
		})
	}
} 