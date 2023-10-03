package main
import "fmt"

func main() {
    var i, n int
    array := [5] int {1,2,3,4,5}
    slice := [] int {}
    
    for i = 0; i < 5; i++ {
        n = array[i]
        if n% 3 == 0 {
            slice = append(slice, n)
        }
    }
    
    fmt.Printf("%d", slice)
}
