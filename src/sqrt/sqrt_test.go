package main

import "testing"

func TestSqrt(t *testing.T) {
	res := sqrt()
	if res != "Code.education Rocks!" {
		t.Error("Erro! correto é Code.education Rocks!", res)
	}
}
