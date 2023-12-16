package main

import (
	"math/rand"
	"fmt"
)

type Problem struct {
	problem string
	solution string
}

func (p *Problem) checkSolution(solution string) bool {
	return solution == p.solution
}

func additionProblem(x int, b int) *Problem {
	var signB string
	var absB int
	if b < 0 {
		signB = "-"
		absB = -b
	} else {
		signB = "+"
		absB = b
	}

	return &Problem{
		problem: fmt.Sprintf("x %s %d = %d", signB, absB, x+b),
		solution: fmt.Sprintf("%d", x),
	}
}

func randomAdditionProblem(allowNegative bool) *Problem {
	x := rand.Intn(21)
	var b int
	if allowNegative {
		b = rand.Intn(41)-20
	} else
	{
		b = rand.Intn(21)
	}
	return additionProblem(x, b)
}