package main
import "fmt"

func seg(slice[]int, c int) {
    for i := 0; i < c-1; i++ {
        for j := i+1; j < c; j++ {
            if slice[i] > slice[j]{
                troca := slice[i]
                slice[i] = slice[j]
                slice[j] = troca
            }
        }
    }
    
    fmt.Printf("Segundo maior valor: %d", slice[c-2])
}

func main() {
    var c, n int
    slice := [] int {}
    
    for {
        fmt.Printf("Digite um valor: ")
        fmt.Scan(&n)
        
        if n == 0 {
            break
        }
        
        c++
        slice = append(slice, n)
    }
    
    seg(slice, c)
}
