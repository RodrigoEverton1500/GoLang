package main
import (
    "fmt" 
)

func f(n int, s[]int) bool {
    l := len(s)
    for i := 0; i < l; i++ {
        if n == s[i] {
            return true
        } else {
            return false
        }
    }
    panic(nil)
}

func main() {
    var n int
    s := [] int {}
    
    for {
        fmt.Scanln(&n)
        if n == 0 {
            break
        }
        s = append(s, n)
    }
    
    fmt.Scanln(&n)
    fmt.Println(f(n, s))
}
