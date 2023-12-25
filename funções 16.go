package main
import (
    "fmt"
    "strings"
)

func f(s string) string {
    strings.ReplaceAll(s, "a", "*")
    strings.ReplaceAll(s, "e", "*")
    strings.ReplaceAll(s, "i", "*")
    strings.ReplaceAll(s, "o", "*")
    strings.ReplaceAll(s, "u", "*")
    
    return s
}

func main() {
    var s string
    
    fmt.Scan(&s)
    fmt.Print(f(s))
}
