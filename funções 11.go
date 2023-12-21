package main
import "fmt"

func f(n int) bool {
    c := 0
    var erro error = nil
    for i := 1; i < n; i++ {
        if n % i == 0 {
            c++
        }
    }
    if c == 1 {
        return true
    } else if c != 1 && n > 0 {
        return false
    }
    panic(erro)
}

func main() {
    n := 0
    fmt.Scan(&n)
    fmt.Println(f(n))
}
