package main
import "fmt"

func main() {
    var i, soma int
    lista := [3]int {1, 2, 3}
    
    for i = 0; i < 3; i++ {
        soma += lista[i]
    }
    
    fmt.Printf("%d\n", lista)
    fmt.Printf("%d", soma)
}
