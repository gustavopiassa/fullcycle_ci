package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(1, 2)

	if total != 3 {
		t.Errorf("Soma deve retornar %d, mas retornou %d", 3, total)
	}
}
