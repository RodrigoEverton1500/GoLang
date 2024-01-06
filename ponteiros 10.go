package main
import (
    "fmt"
    )

func ponteiro(s *[]int, l int) {
    var c int
    var c1 int
    
    for i := 1; i < 100; i++ {
        for j := 1; j < 100; j++ {
            if i % j == 0 {
                c ++
            }
        }
        if c == 2 && c1 < l {
            *s = append(*s, i)
            c1++
        }
        c = 0
    }
}

func main() {
    s := [] int {}
    l := 10
    
    fmt.Println(s)
    ponteiro(&s, l)
    fmt.Println(s)
}
