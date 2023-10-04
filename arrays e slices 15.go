package main
import "fmt"

func main() {
    array := [10] float32 {1,2,3,4,5,6,7,8,9,10}
    slice := [] float32 {}
    
    for i := 0; i < 10; i++ {
        if array[i] > 5 {
            slice = append(slice, array[i])
        }
    }
    
    fmt.Printf("%f", slice)
}
