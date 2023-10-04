package main
import "fmt"

func main() {
    var i, j, leng, c1, c2 int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&leng)
    
    lista := [] int {}
    
    for i = 1; i > -1; i++ {
        if c2 == leng {
            break
        }
        for j = 1; j <= i; j++ {
            if i % j == 0 {
                c1++
            }
        }
        if c1 == 2 {
            c2++
            lista = append(lista, i)
        }
        c1 = 0
    }
    
    fmt.Printf("%d", lista)
}
