package main
import "fmt"

func main() {
    produto := 1
    lista := [4]int {1, 2, 3, 4}
    
    for i := 0; i < 4; i++ {
        produto *= lista[i]
    }
    
    fmt.Printf("Produto dos elementos: %d", produto)
}
