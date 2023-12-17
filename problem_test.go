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
		{"Correct solution with different format", Problem{"x + 1 = 2", "1"}, " + 1 ", true},
		{"Wrong solution", Problem{"x + 1 = 2", "1"}, "0", false},
	}

	for _,tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			check := tc.p.checkSolution(tc.s)
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
				t.Errorf(`additionProblem(%d, %d) = Problem{"%s", "%s"}, want Problem{"%s", "%s"}`, tc.x, tc.b, p.problem, p.solution, tc.pe.problem, tc.pe.solution)
			}
		})
	}
}

func TestMxbProblem(t *testing.T) {
	var testCases = []struct{
		name string
		x int
		m int
		b int
		want Problem
		wantErr bool
	} {
		{
			name: "Positive m",
			x: 2,
		 	m: 3,
		 	b: 10,
		 	want: Problem{"3x + 10 = 16", "2"},
			wantErr: false,
		},
		{
			name: "Negative m",
			x: 2,
			m: -3,
			b: 10,
			want: Problem{"-3x + 10 = 4", "2"},
			wantErr: false,
		},
		{
			name: "Zero m",
			x: 2,
			m: 0,
			b: 10,
			want: Problem{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func (t *testing.T) {
			p, err := mxbProblem(tc.x, tc.m, tc.b)
			if tc.wantErr && err == nil {
				t.Errorf(`mxbProblem(%d, %d, %d) expected error`, tc.x, tc.m, tc.b)
			}
			if !tc.wantErr && err != nil {
				t.Errorf(`mxbProblem(%d, %d, %d) got unexpected error %v`, tc.x, tc.m, tc.b, err)
			}
			if err == nil && (p.problem != tc.want.problem || !p.checkSolution(tc.want.solution)) {
				t.Errorf(`mxbProblem(%d, %d, %d) = {Problem{"%s", "%s"}, nil}, want {Problem{"%s", "%s"}, nil}`, tc.x, tc.m, tc.b, p.problem, p.solution, tc.want.problem, tc.want.solution)
			}

		})
	}
}