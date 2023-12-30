package main
import (
    "fmt"
    "math"
    )

type circulo struct {
    raio float32
}

func f(class circulo) {
    n := class.raio
    n *= class.raio
    n *= math.Pi
    fmt.Println(n)
}

func main() {
    var n float32
    var u circulo
    
    fmt.Scanln(&n)
    u.raio = n
    
    f(u)
}
