package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	x := "1 + 1 + 1"

	t.Logf("Evaluate %d", Evaluate(x))

	if Evaluate(x) != 3 {
		t.Errorf("SALAH! harusnya 3")
	}
}

func TestMinus(t *testing.T) {
	x := "10 - 1 - 3"

	t.Logf("Evaluate %d", Evaluate(x))

	if Evaluate(x) != 6 {
		t.Errorf("SALAH! harusnya 6")
	}
}

func TestRandom(t *testing.T) {
	x := "100 - 50 + 3"

	t.Logf("Evaluate %d", Evaluate(x))

	if Evaluate(x) != 53 {
		t.Errorf("SALAH! harusnya 53")
	}
}
