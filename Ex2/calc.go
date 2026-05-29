package main

import (
	"errors"
	"fmt"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro")
		}
		return a / b, nil
	}
	return 0, errors.New("opération inconnue")
}

func creerOperation(op string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		r, _ := operer(a, b, op)
		return r
	}
}

func main() {
	for {
		var a, b float64
		var op string
		fmt.Print("> ")
		fmt.Scan(&a, &b, &op)

		if op == "quit" {
			break
		}

		r, err := operer(a, b, op)
		if err != nil {
			fmt.Println("erreur :", err)
		} else {
			fmt.Println("=", r)
		}
	}
}
