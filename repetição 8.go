package main
import "fmt"

func main() {
    var i, n uint
    
    fmt.Printf("Digite um número inteiro positivo: ")
    fmt.Scan(&n)
    
    for i = 1; i < n; i++ {
        if n%i == 0 {
            fmt.Printf("\n%d / %d", n, i)
        }
    }
}
