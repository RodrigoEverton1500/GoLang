package main
import (
    "fmt"
    )

func ponteiro(p *int) {
    *p = 5
}

func main() {
    var n int = 10
    
    fmt.Println(n)
    ponteiro(&n)
    fmt.Println(n)
}
