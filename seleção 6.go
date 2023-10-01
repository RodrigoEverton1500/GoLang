package main
import "fmt"

func main() {
    var n1, n2 int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n1)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n2)
    
    if n1 < 0 || n2 < 0 {
        fmt.Printf("Soma: %d", n1+n2)
    } else if n1 > 0 && n2 > 0 {
        fmt.Printf("Multiplicação: %d", n1*n2)
    } else {
        fmt.Printf("Valores inválidos")
    }
}
