package main
import (
    "fmt"
    )

func ponteiro(n *int) {
    if *n % 2 == 0 {
        *n = 0
    } else {
        *n = 1
    }
}

func main() {
    var n int = 9
    
    fmt.Println(n)
    ponteiro(&n)
    fmt.Println(n)
}
