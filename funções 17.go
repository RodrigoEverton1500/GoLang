package main
import (
    "fmt"
    "sort"
    )
func f(s[]string) string{
    sort.Strings(s)
    var r string
    
    for _, i := range s {
        r += i + " "
    }
    
    return r
}

func main() {
    var str string
    s := [] string {}
    
    for {
        fmt.Scanln(&str)
        if str == "0" {
            break
        }
        s = append(s, str)
    }
    
    fmt.Println(f(s))
}
