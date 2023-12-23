package main
import (
    "fmt" 
    "strings"
)

func f(s []string) string {
    l := len(s)
    var upper string
    var r string
    
    for i := 0; i < l; i++ {
        upper = strings.Title(s[i])
        if upper == s[i] {
            r += s[i]
        }
    }
    
    return r
}

func main() {
    s1 := ""
    s2 := [] string {}
    
    for {
        fmt.Scanln(&s1)
        if s1 == "0" {
            break
        }
        s2 = append(s2, s1)
    }
    
    fmt.Println(f(s2))
}
