package main

import (
	"fmt"
)

func main() {
	var (
		numero1 float64
		numero2 float64
		funcion string
	)

	fmt.Println("Hola! Esta es una calculadora basica en go.")

	fmt.Println("Escriba el primer numero: ")
	fmt.Scanf("%f", &numero1)

	fmt.Println("Escriba el segundo numero: ")
	fmt.Scanf("%f", &numero2)

	fmt.Println("Que funcion deseas realizar?= (+, -, /, *)")
	fmt.Scanf("%s", &funcion)

	fmt.Println("El resultado es: ", resultado(numero1, numero2, funcion))
}

func resultado(a, b float64, f string) float64 {
	var result float64
	switch f {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	default:
		fmt.Println("Esa operacion no se puede realizar")
	}

	return result
}
