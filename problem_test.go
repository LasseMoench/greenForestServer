package main

import (
	"testing"
)

func TestCheckSolution(t *testing.T) {
	var testCases = []struct{
		name string
		p Problem
		s string
		want bool
	}{
		{"Correct solution", Problem{"x + 1 = 2", "1"}, "1", true},
		{"Wrong solution", Problem{"x + 1 = 2", "1"}, "0", false},
	}

	for _,tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			check : = tc.p.checkSolution(tc.s)
			if check != tc.want {
				t.Errorf(`Problem{"%s", "%s"}.checkSolution("%s") = %t, want %t`, tc.p.problem, tc.p.solution, tc.s, check, tc.want)
			}
		})
	}
}

func TestAdditionProblem(t *testing.T) {
	var testCases = []struct{
		name string
		x int
		b int
		pe Problem
	}{
		{"Positive b", 5, 3, Problem{"x + 3 = 8", "5"}},
		{"Negative b", 5, -3, Problem{"x - 3 = 2", "5"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func (t *testing.T){
			p := additionProblem(tc.x, tc.b)
			if p.problem != tc.pe.problem || !p.checkSolution(tc.pe.solution) {
				t.Errorf(`additionProblem(%d, %d) == Problem{"%s", "%s"}, want Problem{"%s", "%s"}`, tc.x, tc.b, p.problem, p.solution, tc.pe.problem, tc.pe.solution)
			}
		})
	}
}