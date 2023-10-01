package main
import "fmt"

func main() {
    var n1, n2 int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n1)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n2)
    
    if n1 > n2 {
        fmt.Printf("\n%d é o maior", n1)
    } else if n2 > n1 {
        fmt.Printf("\n%d é o maior", n2)
    } else {
        fmt.Printf("\n%d é igual a %d", n1, n2)
    }
}
