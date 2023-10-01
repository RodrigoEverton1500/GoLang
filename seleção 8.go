package main
import "fmt"

func main() {
    var dia int
    
    fmt.Printf("Digite o dia da semana de 1 a 7: ")
    fmt.Scan(&dia)
    
    switch (dia) {
    case 1: fmt.Printf("Domingo")
    case 2: fmt.Printf("Segunda")
    case 3: fmt.Printf("Terça")
    case 4: fmt.Printf("Quarta")
    case 5: fmt.Printf("Quinta")
    case 6: fmt.Printf("Sexta")
    case 7: fmt.Printf("Sábado")
    }
}
