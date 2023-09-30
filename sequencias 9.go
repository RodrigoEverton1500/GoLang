package main
import "fmt"

func main() {
    var preco float32
    
    fmt.Printf("Digite preço do produto: ")
    fmt.Scan(&preco)
    
    fmt.Printf("Preço com desconto de 10%%: R$%.2f", preco*0.9)
}
