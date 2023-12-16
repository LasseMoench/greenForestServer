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

func additionProblem() *Problem {
	x := rand.Intn(21)
	b := rand.Intn(21)

	return &Problem{
		problem: fmt.Sprintf("x + %d = %d", b, x+b),
		solution: fmt.Sprintf("%d", x),
	}
}