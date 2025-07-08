package calculo

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(2, 3)
	esperado := 5
	if resultado != esperado {
		t.Errorf("esperado %d, mas obteve %d", esperado, resultado)
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(2, 3)
	}
} 