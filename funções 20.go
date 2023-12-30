package main
import (
    "fmt"
    "strings"
    )

func f(s []string) []string {
    r := [] string {}
    var n int
    
    for _, str := range s {
        n = strings.Count(str, "")
        if n > 5 {
            r = append(r, str)
        }
    }
    
    return r
}

func main() {
    s := [] string {}
    var str string
    
    for {
        fmt.Scanln(&str)
        if str == "0" {
            break
        }
        s = append(s, str)
    }
    
    fmt.Print(f(s))
}
