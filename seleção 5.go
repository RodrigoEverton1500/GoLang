package main
import "fmt"

func main() {
    var n int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n)
    
    if n%3 == 0 && n%5 ==0 {
        fmt.Printf("%d é multiplo de 3 e 5", n)
    } else {
        fmt.Printf("%d não é multiplo de 3 e 5", n)
    }
}
