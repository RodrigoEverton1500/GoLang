package main
import "fmt"

func main() {
    var soma, num int
    
    for i := 0; i < 3; i++ {
        fmt.Printf("\nDigite um valor: ")
        fmt.Scan(&num)

        soma += num
    }
    
    fmt.Printf("\nSoma dos valores: %d", soma)
}
