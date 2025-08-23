package main

import (
"fmt"
"askbow.io/mymath"
)


func main() {
    a := 146
    b := 2
    fmt.Printf("a = %d\tb = %d\n", a, b)
    fmt.Printf("Сумма: %d\n", mymath.Add(a, b))
    fmt.Printf("Произведение: %d\n", mymath.Multiply(a, b))
}