package main
import "fmt"

func main() {
    var n int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n)
    
    if n%2 == 0 {
        fmt.Printf("\n%d é par", n)
    } else {
        fmt.Printf("\n%d é impar", n)
    }
}
