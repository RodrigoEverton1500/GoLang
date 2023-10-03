package main
import "fmt"

func main() {
    var n float32
    var i int
    lista := [6] float32 {}
    
    fmt.Printf("Digite número real: ")
    fmt.Scan(&n)
    
    for i = 0; i < 6; i++ {
        lista[i] = n
    }
    
    fmt.Printf("%.3f", lista)
}
