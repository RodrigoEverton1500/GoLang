package main
import "fmt"

func main() {
    var diaria, salario float32
    
    fmt.Printf("Digite dias trabalhados: ")
    fmt.Scan(&diaria)
    
    fmt.Printf("Digite valor da diária: ")
    fmt.Scan(&salario)
    
    fmt.Printf("Salário: %.2f", salario*diaria)
}
