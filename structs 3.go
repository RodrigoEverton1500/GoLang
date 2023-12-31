package main
import (
    "fmt"
    )

type Triangulo struct {
    base float32
    h float32
}

func f(class Triangulo) {
    area := class.base * class.h
    area /= 2
    
    fmt.Print(area)
}

func main() {
    var u Triangulo
    
    fmt.Print("Base: ")
    fmt.Scan(&u.base)
    fmt.Print("Altura: ")
    fmt.Scan(&u.h)
    
    f(u)
}
