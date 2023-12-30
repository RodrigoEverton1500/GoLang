package main
import "fmt"

func f(n int) []int {
    var c int
    s := [] int {}
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if i % j == 0 {
                c++
            }
        }
        if c == 2 {
            s = append(s, i)
        }
        c = 0
    }
    
    return s
    panic(nil)
}

func main() {
    var n int
    
    fmt.Scanln(&n)
    fmt.Print(f(n))
}
