package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mult(a, b float64) float64 {
	return a * b
}

func devide(a, b float64) float64 {
	if b != 0 {
		return a / b
	} else {
		fmt.Println("Division Error : Division par zéro")
		return 0
	}
}

func main() {
	var num1, num2 float64
	var operator string
	var result float64

	fmt.Print("Entrez le premier nombre : ")
	fmt.Scanln(&num1)

	fmt.Print("Entrez l'opérateur (+, -, *, /) : ")
	fmt.Scanln(&operator)

	fmt.Print("Entrez le deuxième nombre : ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = sub(num1, num2)
	case "*":
		result = mult(num1, num2)
	case "/":
		result = devide(num1, num2)
	default:
		fmt.Println("Opérateur invalide")
		return
	}

	fmt.Printf("Résultat : %.2f\n", result)
	fmt.Println("Appuie sur Entrée pour quitter...")
	fmt.Scanln()
}
