package main
import (
    "fmt"
    )

func main() {
    var p *int
    var n int = 10
    var soma int
    p = &n
    
    fmt.Println(*p, n)
    
    for i := 0; i <= n; i++ {
        soma += i
    }
    n = soma
    
    fmt.Println(*p, n)
}
