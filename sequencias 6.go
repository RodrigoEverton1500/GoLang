package main
import "fmt"

func main() {
    var salario float32
    
    fmt.Printf("Digite salário: ")
    fmt.Scan(&salario)
    
    fmt.Printf("Novo salário: %.2f", salario*1.15)
}
