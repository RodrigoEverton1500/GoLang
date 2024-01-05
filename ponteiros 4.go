package main
import (
    "fmt"
    "strconv"
    )

func ponteiro(n *int64) {
    s := strconv.FormatInt(int64(*n), 10)
    l := len(s)
    *n = int64(s[l-2]) - 96 + int64(s[l-1])
}

func main() {
    var n int64 = 4321
    
    fmt.Println(n)
    ponteiro(&n)
    fmt.Println(n)
}
