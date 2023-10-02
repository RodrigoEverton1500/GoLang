package main
import "fmt"

func main() {
    var i, j, c, n, troca int
    lista := []int {}
    
    for {
        fmt.Printf("Digite um valor: ")
        fmt.Scan(&n)
        
        if n == 0 {
            break
        }
        
        c++
        lista = append(lista, n)
    }
    
    for i = 0; i < c-1; i++ {
        for j = i+1; j < c; j++{
            if lista[i] > lista[j] {
                troca = lista[i]
                lista[i] = lista[j]
                lista[j] = troca
            }
        }
    }
    
    for i = 0; i < c; i++ {
        fmt.Printf("\n%d", lista[i])
    }
}
