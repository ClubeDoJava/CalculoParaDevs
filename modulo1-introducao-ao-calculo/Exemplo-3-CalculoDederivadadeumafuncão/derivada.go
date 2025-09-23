package main

import (
    "fmt"
    "math"
)

// Função que queremos calcular a derivada
func f(x float64) float64 {
    return math.Pow(x, 2)
}

// Função que calcula a derivada usando a definição de limite
func derivative(f func(float64) float64, x float64) float64 {
    h := 1e-10
    return (f(x+h) - f(x)) / h
}

func main() {
    x := 2.0
    fmt.Printf("A derivada de f(x) em x = %v é aproximadamente %v\n", x, derivative(f, x))
}
