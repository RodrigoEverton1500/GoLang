package main
import "fmt"

func main() {
    var i, n, b int
    lista := [] int {1, 2, 3, 4, 5}
    
    fmt.Printf("Adicione um valor a lista: ")
    fmt.Scan(&n)
    
    for i = 0; i < 5; i++ {
        if lista[i] == n {
            b = 1
        }
    }
    
    if b == 0 {
        lista = append(lista, n)
    }
    
    fmt.Printf("%d", lista)
}
