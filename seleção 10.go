package main
import "fmt"

func main() {
    var idade int
    
    fmt.Printf("Digite idade: ")
    fmt.Scan(&idade)
    
    if idade < 10 && idade >= 0 {
        fmt.Printf("%d - Mirim", idade)
    } else if idade < 14 && idade > 9 {
        fmt.Printf("%d - Infantil", idade)
    } else if idade < 18 && idade > 13 {
        fmt.Printf("%d - Juvenil", idade)
    } else if idade > 17 {
        fmt.Printf("%d - Adulto", idade)
    }
}
