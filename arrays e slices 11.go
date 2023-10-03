package main
import "fmt"

func main() {
    var i, j, troca int
    lista := [] int {10,9,8,7,6,5,4,3,2,1}
    
    for i = 0; i < 9; i++ {
        for j = i+1; j < 10; j++ {
            if lista[i] > lista[j] {
                troca = lista[i]
                lista[i] = lista[j]
                lista[j] = troca
            }
        } 
    }
    
    fmt.Printf("Maior valor: %d\nMenor valor: %d", lista[9], lista[0])
}
