package main
import "fmt"

func main() {
    var n1, n2, n3, media float32
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n1)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n2)
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n3)
    
    media = (n1*2) + (n2*3) + (n3*5)
    media = media / 3
    
    fmt.Printf("Média: %2.2f", media)
}
