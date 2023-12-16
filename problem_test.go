package main

import (
	"testing"
	"fmt"
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

func TestAdditionProblem(t *testing.T) {
	x := 5
	p := additionProblem(x, -3)

	if !p.checkSolution(fmt.Sprintf("%d", x)){
		t.Fatalf(`Problem {'%s', '%s'} did not accept '%d' as the solution`, p.problem, p.solution, x)
	}
}