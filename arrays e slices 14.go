package main
import "fmt"

func main() {
    var troca, n1, n2 int
    lista := [] int {1,2,3,4,5,6,7,8}
    
    fmt.Printf("Digite o primeiro índice: ")
    fmt.Scan(&n1)
    
    fmt.Printf("Digite o segundo índice: ")
    fmt.Scan(&n2)
    
    troca = lista[n1]
    lista[n1] = lista[n2]
    lista[n2] = troca
    
    fmt.Printf("%d", lista)
}
