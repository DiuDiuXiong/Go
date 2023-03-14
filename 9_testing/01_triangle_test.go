package main

import "testing"

/*
*
1. store test data in struct
2. testing function should have (t *testing.T) as input, then whenever something went wrong, use t.Errorf() to report error without blocking others
3. Rather than use idea, go to terminal ==> go to test file directory ==> `go test .` to run tests
*/
func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		//{1, 1, 100},
		//{2, 2, 200},
		//{3, 3, 300},
		{300000, 400000, 500000},
	}

	for _, tt := range tests {
		if actual := CalcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("CalcTriangle(%d, %d): "+"got %d, expect %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
