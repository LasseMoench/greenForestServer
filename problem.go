package main

import (
	"math/rand"
	"fmt"
	"strconv"
	"strings"
	"errors"
)

type Problem struct {
	problem string
	solution string
}

func (p *Problem) checkSolution(solution string) bool {
	pis, err := strconv.Atoi(p.solution)
	if err != nil { panic(fmt.Sprintf("Invalid solution '%s'", p.solution)) }

	r := strings.NewReplacer(" ", "", "\t", "", "\n", "");
	is, err := strconv.Atoi(r.Replace(solution))
	if err != nil {
		 return false
	}
	return is == pis
}

func ProblemForClient(c *Client) *Problem {
	if c.Points > 100 {
		return randomMxbProblem()
	} else if c.Points > 50 {
		return randomAdditionProblem(true)
	} else {
		return randomAdditionProblem(false)
	}
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

func mxbProblem(x int, m int, b int) (*Problem, error) {
	if m == 0 {
		return &Problem{}, errors.New("Cannot have m==0 in mx + b problem")
	}
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
		problem: fmt.Sprintf("%dx %s %d = %d", m, signB, absB, m * x + b),
		solution: fmt.Sprintf("%d", x),
	}, nil
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

func randomMxbProblem() *Problem {
	x := rand.Intn(21)
	m := rand.Intn(9) + 1
	b := rand.Intn(100) - 50

	p, err := mxbProblem(x, m, b)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return p
}