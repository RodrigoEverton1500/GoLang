package main
import "fmt"

func main() {
    var soma int
    array := [10] int {1,2,3,4,5,6,7,8,9,10}
    
    for i := 0; i < 10; i++ {
        if array[i] % 2 == 0 {
            soma += array[i]
        }
    }
    
    fmt.Printf("%d", soma)
}
