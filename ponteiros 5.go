package main
import (
    "fmt"
    "math"
    )

func ponteiro(n *float32) {
    *n += math.Pi
    *n /= 2
}

func main() {
    var n float32 = 10
    
    fmt.Println(n)
    ponteiro(&n)
    fmt.Println(n)
}
