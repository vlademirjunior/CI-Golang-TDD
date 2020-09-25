package main

import "testing"

func TestSoma(t *testing.T) {
	got := soma(10, 20)

	if got != 30 {
		t.Errorf("Teste = %d; want 30", got)
	}
}
