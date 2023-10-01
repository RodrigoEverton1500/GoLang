package main
import "fmt"

func main() {
    var i, n, c, soma int
    lista := []int {}
    
    for i = 0; i > -1; i++ {
        fmt.Printf("Digite um valor: ")
        fmt.Scan(&n)
        
        if n == 0 {
            break
        }
        
        c += 1
        
        lista = append(lista, n)
    }
    
    for i = 0; i < c; i++ {
        soma += lista[i]
    }
    
    fmt.Printf("\nSoma = %d\nMédia = %d", soma, soma/c)
}
