package main
import "fmt"

func main() {
    var n1, n2, n3 int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n1)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n2)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n3)
    
    if n1 > (n2+n3)/2 {
        fmt.Printf("\n%d é o maior valor", n1)
    } else if n2 > (n1+n3)/2 {
        fmt.Printf("\n%d é o maior valor", n2)
    } else {
        fmt.Printf("\n%d é o maior valor", n3)
    }
}
