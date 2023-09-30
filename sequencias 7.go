package main
import "fmt"

func main() {
    var n int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n)
    
    fmt.Printf("\nAntecessor: %d", n-1)
    fmt.Printf("\nSucessor: %d", n+1)
}
