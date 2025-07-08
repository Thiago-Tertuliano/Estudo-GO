package main

import "testing"

func TestSomaSimples(t *testing.T) {
	esperado := 4
	resultado := 2 + 2
	if resultado != esperado {
		t.Errorf("esperado %d, obteve %d", esperado, resultado)
	}
} 