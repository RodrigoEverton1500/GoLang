package main
import "fmt"

func main() {
    var peso, altura, imc, sexo float32
    
    fmt.Printf("Digite o peso: ")
    fmt.Scan(&peso)
    
    fmt.Printf("Digite o altura: ")
    fmt.Scan(&altura)
    
    fmt.Printf("Digite 1 para masculino e 2 para feminino: ")
    fmt.Scan(&sexo)
    
    imc = peso / (altura*altura)
    fmt.Printf("IMC = %.2f\n", imc)
    
    if sexo == 1 {
        if imc >= 25 {
            fmt.Printf("Acima do peso")
        } else if imc >= 20 && imc <= 24.99 {
            fmt.Printf("Dentro do peso")
        } else {
            fmt.Printf("Abaixo do peso")
        }
    } else if sexo == 2 {
        if imc >= 24 {
            fmt.Printf("Acima do peso")
        } else if imc >= 19 && imc <= 23.99 {
            fmt.Printf("Dentro do peso")
        } else {
            fmt.Printf("Abaixo do peso")
        }
    }
}
