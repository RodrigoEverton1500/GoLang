package main
import "fmt"

func main() {
    lista := []int {1, 2, 3, 4, 5}
    
    lista = append(lista[:3], lista[4:]...)
    
    fmt.Printf("%d\n", lista)
}
