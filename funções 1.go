package main
import "fmt"

func media(slice[]int, leng int) {
    var soma, i int
    
    for i = 0; i < leng; i++ {
        soma += slice[i]
    }
    
    soma = soma / leng
    fmt.Printf("Média aritimética: %d", soma)
}

func main() {
    var i, n, leng int
    slice := [] int {}
    
    fmt.Printf("Digite o tamanho do slice: ")
    fmt.Scan(&leng)
    
    for i = 0; i < leng; i++ {
        fmt.Printf("[%d]: ", i)
        fmt.Scan(&n)
        
        slice = append(slice, n)
    }
    
    media(slice, leng)
}
