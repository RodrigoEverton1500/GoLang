package main
import "fmt"

func main() {
    var sal float32
    
    fmt.Printf("Digite o salário: ")
    fmt.Scan(&sal)
    
    if sal < 1000 {
        fmt.Printf("Novo salário: %.2f", sal*1.1)
    } else if sal >= 1000 {
        fmt.Printf("Novo salário: %.2f", sal*1.05)
    }
}
