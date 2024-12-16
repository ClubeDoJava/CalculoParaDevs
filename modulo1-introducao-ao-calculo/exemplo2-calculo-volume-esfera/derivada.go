package main

import (
    "fmt"
    "math"
)

func f(x float64) float64 {
    return math.Pow(x, 2)
}

func derivative(f func(float64) float64, x float64) float64 {
    h := 1e-10
    return (f(x+h) - f(x)) / h
}

func main() {
    x := 2.0
    fmt.Printf("A derivada de f(x) em x = %v é aproximadamente %v\n", x, derivative(f, x))
}
