package main
import "fmt"

func main() {
    var i, n int
    lista := [10] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Printf("Digite um valor para ser buscado: ")
    fmt.Scan(&n)
    
    for i = 0; i < 10; i++ {
        if lista[i] == n {
            fmt.Printf("%d encontrado na posição: [%d]", n, i)
        }
    }
}
