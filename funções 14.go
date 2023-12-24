package main
import (
    "fmt" 
)

func f(s1[]int, s2[]int) []int {
    r := [] int {}
    l1 := len(s1)
    l2 := len(s2)
    
    for i := 0; i < l1; i++ {
        for j := 0; j < l2; j++ {
            if s1[i] == s2[j] {
                r = append(r, s2[j])
            }
        }
    }
    
    return r
}

func main() {
    var n int
    s1 := [] int {}
    s2 := [] int {}
    
    for {
        fmt.Scanln(&n)
        if n == 0 {
            break
        }
        s1 = append(s1, n)
    }
    
    for {
        fmt.Scanln(&n)
        if n == 0 {
            break
        }
        s2 = append(s2, n)
    }
    
    fmt.Println(f(s1, s2))
}
