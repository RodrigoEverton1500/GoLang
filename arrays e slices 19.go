package main
import "fmt"

func main() {
    var i int
    array1 := [5] int {1,2,3,4,5}
    array2 := [5] int {6,7,8,9,10}
    array3 := [5] int {}
    
    for i = 0; i < 5; i++ {
        array3[i] = array1[i] + array2[i]
    }
    
    fmt.Printf("%d", array3)
}
