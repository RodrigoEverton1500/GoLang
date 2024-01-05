package main
import (
    "fmt"
    )

func ponteiro(n *int) {
    l := *n
    *n = 0
    for i := 1; i <= l; i++ {
        *n += i
    }
}

func main() {
    var n int = 5
    
    fmt.Println(n)
    ponteiro(&n)
    fmt.Println(n)
}
