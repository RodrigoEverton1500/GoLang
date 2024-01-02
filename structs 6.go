package main
import "fmt"

type Funcionario struct {
    nome string
    salario float32
    idade int
}

func porcentagem(class Funcionario) {
    var n float32
    
    fmt.Print("Porcentagem: ")
    fmt.Scan(&n)
    
     n = (class.salario/100) * (n+100)
     fmt.Printf("Novo salário: %.2f\n", n)
}

func servico(class Funcionario) {
    var n int
    
    n = class.idade - 18
    fmt.Printf("Tempo de serviço: %d", n)
}

func main() {
    var funcionario Funcionario
    
    fmt.Print("Nome: ")
    fmt.Scan(&funcionario.nome)
    fmt.Print("Salário: ")
    fmt.Scan(&funcionario.salario)
    fmt.Print("Idade: ")
    fmt.Scan(&funcionario.idade)
    
    porcentagem(funcionario)
    servico(funcionario)
}
