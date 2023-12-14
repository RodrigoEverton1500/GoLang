package main
import "fmt"

func f() {
    n := 0
    c := 0
    s := [100] int {}
    
    for i := 0; i < 100; i++ {
        fmt.Scan(&n)
        if n == 0 {
            break
        }
        
        s[i] = n
        c++
    }
    
    for i := 0; i < c; i++ {
        n := s[i]
        fmt.Println(parametro(n))
    }
}

func parametro(n int) int {
    n *= 5
    return n
}

func main() {
    f()
}
