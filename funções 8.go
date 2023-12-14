package main
import "fmt"

func f() []int {
    var i, c, n int = 0, 0, 0
    s := [] int {}
    r := [] int {}
    
    for {
        fmt.Scan(&n)
        if n == 0 {
            break
        }
        s = append(s, n)
        c++
        i++
    }
    
    for i := 0; i < c; i++ {
        if s[i] % 2 == 0 {
            r = append(r, s[i])
        }
    }
    
    return r
}

func main() {
    fn := f()
    fmt.Println(fn)
}
