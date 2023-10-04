package main
import "fmt"

func main() {
    array := [10] int {1,2,3,4,5,6,7,8,9,10}
    slice := [] int {}
    
    for i := 0; i < 10; i++ {
        if array[i] % 2 == 0 {
            slice = append(slice, array[i])
        }
    }
    
    fmt.Printf("%d", slice)
}
