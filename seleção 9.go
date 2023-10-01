package main
import "fmt"

func main() {
    var n1, n2, n3 float32
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n1)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n2)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n3)
    
    if n1 > n2 && n2 > n3 {
        fmt.Printf("%f > %f > %f", n1, n2 ,n3)
    } else if n2 > n1 && n1 > n3 {
        fmt.Printf("%f > %f > %f", n2, n1 ,n3)
    } else if n3 > n2 && n2 > n1 {
        fmt.Printf("%f > %f > %f", n3, n2 ,n1)
    } else if n3 > n1 && n1 > n2 {
        fmt.Printf("%f > %f > %f", n3, n1 ,n2)
    }
}
