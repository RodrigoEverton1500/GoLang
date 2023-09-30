package main
import "fmt"

func main() {
    var idade, dias int
    
    fmt.Printf("Digite idade: ")
    fmt.Scan(&idade)
    
    dias = 365 * idade
    dias += (idade/4)
    
    fmt.Printf("Idade em dias: %d", dias)
}
