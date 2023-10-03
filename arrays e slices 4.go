package main
import "fmt"

func main() {
    var i, soma, n, c int
    
    fmt.Printf("Digite o tamanho do slice: ")
    fmt.Scan(&c)
    lista := []int {}
    
    for i = 0; i < c; i++ {
        lista = append(lista, 0)
    }
    
    for i = 0; i < c; i++{
        fmt.Printf("Digite um valor: ")
        fmt.Scan(&n)
        
        soma += n
        lista[i] = n
    }
    
    fmt.Printf("%d\n", lista)
    fmt.Printf("Soma dos elementos: %d", soma)
}
