package main
import "fmt"

func main() {
    var i, n int
    
    fmt.Printf("Digite um número: ")
    fmt.Scan(&n)
    
    for i = 1; i < 11; i++ {
        fmt.Printf("\n%d X %2d = %5d", n, i, n*i)
    }
}
