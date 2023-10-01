package main
import "fmt"

func main() {
    var i int
    
    for i = 0; i < 31; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
}
