package main

import (
	"testing"
)

func TestCheckSolution_WrongInput(t *testing.T) {
	p := Problem{
		problem: "x + 9 = 11",
		solution: "2",
	}

	givenSolution := "3"
	expected := false
	check := p.checkSolution(givenSolution)
	if check != expected {
		t.Fatalf(`p.checkSolution("%s") = %t, want %t`, givenSolution, check, expected )
	}
}

func TestCheckSolution_CorrectInput(t *testing.T) {
	p := Problem{
		problem: "x + 9 = 11",
		solution: "2",
	}

	givenSolution := "2"
	expected := true
	check := p.checkSolution(givenSolution)
	if check != expected {
		t.Fatalf(`p.checkSolution("%s") = %t, want %t`, givenSolution, check, expected )
	}
}

func TestEqAdd(t *testing.T) {
	p := additionProblem(true)

	if !p.checkSolution("p.solution"){
		t.Fatalf(`Problem '%s': Solution '%s' was not accepted`, p.problem, p.solution)
	}
}