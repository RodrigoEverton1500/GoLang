package main
import "fmt"

func f(s[]int) []int {
    l := len(s)
    t := 0
    
    for i := 0; i < l-1; i++ {
        for j := i+1; j < l; j++ {
            if s[i] > s[j] {
                t = s[i]
                s[i] = s[j]
                s[j] = t
            }
        }
    }
    
    return s
}

func main() {
    n := 0
    s := [] int {}
    
    for {
        fmt.Scanln(&n)
        if n == 0 {
            break
        }
        s = append(s, n)
    }
    
    fmt.Println(f(s))
}
