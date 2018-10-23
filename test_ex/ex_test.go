package main

import "testing"

func mul(a, b float64) float64 {
	return a * b
}

func TestMul(t *testing.T) {
	res1 := mul(1.1, 1.1)
	if res1 != 1.21 {
		t.Logf("res1 = %f", res1)
		t.Fail()
	}

	a, b := 0.0, 1.0
	t.Logf("%f * %f", a, b)
	res := mul(a, b)
	if res != 0.0 {
		t.Fail()
	}
}
