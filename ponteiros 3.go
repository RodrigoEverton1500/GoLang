package main
import (
    "fmt"
    )

func ponteiro(s *string) {
    var r string
    for _, i := range *s {
        r = string(i) + r
    }
    *s = r
}

func main() {
    var s string = "Olá td bem?"
    
    fmt.Println(s)
    ponteiro(&s)
    fmt.Println(s)
}
