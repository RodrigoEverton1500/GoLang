package main
import "fmt"

func main() {
    var i, j, b, troca int
    array := [5] int {2,4,6,8,10}
    cresc := [5] int {}

    for i = 0; i < 5; i++ {
        cresc[i] = array[i]
    }

    for i = 0; i < 4; i++ {
        for j = i+1; j < 5; j++ {
            if cresc[i] > cresc[j] {
                troca = cresc[i]
                cresc[i] = cresc[j]
                cresc[j] = troca
            }
        }
    }
    
    for i = 0; i < 5; i++ {
        if array[i] != cresc[i] {
            b++
        }
    }
    
    if b == 0 {
        fmt.Printf("Array em ordem crescente")
    } else {
        fmt.Printf("Array não está em ordem crescente")
    }
}
