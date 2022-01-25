package main

import "testing"

func Soma(a int, b int) int {
	return a + b ;
}

func TestSoma(t *testing.T){

	total := Soma(15,15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalida: Resultado: %d Esperado: %d", total,30)
	}
}
