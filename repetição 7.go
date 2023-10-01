package main
import "fmt"

func main() {
    var i int
    f := "Fizz"; b := "Buzz"
    
    for i = 1; i < 101; i++ {
        if i%3 == 0 {
            fmt.Printf(f)
        } 
        if i%5 == 0 {
            fmt.Printf(b)
        } else if i%3 != 0 && i%5 != 0 {
            fmt.Printf("%d", i)
        }
        fmt.Println()
    }
}
