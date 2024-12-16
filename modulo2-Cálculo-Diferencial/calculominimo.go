package main

import (
    "fmt"
    "math"
)

// Definindo a função
func f(x float64) float64 {
    return math.Pow(x-2, 2)
}

// Função para realizar a busca linear
func linearSearch(f func(float64) float64, a, b, tol float64) float64 {
    // Método da busca linear para encontrar mínimo
    var x1, x2 float64
    goldenRatio := (math.Sqrt(5) - 1) / 2
    for (b - a) > tol {
        x1 = b - (b-a)*goldenRatio
        x2 = a + (b-a)*goldenRatio
        if f(x1) < f(x2) {
            b = x2
        } else {
            a = x1
        }
    }
    return (a + b) / 2
}

func main() {
    a := 0.0   // Limite inferior
    b := 5.0   // Limite superior
    tol := 1e-5 // Tolerância para a convergência

    minPoint := linearSearch(f, a, b, tol)
    minValue := f(minPoint)

    fmt.Printf("O ponto de mínimo ocorre em x = %v\n", minPoint)
    fmt.Printf("O valor mínimo da função é f(x) = %v\n", minValue)
}
