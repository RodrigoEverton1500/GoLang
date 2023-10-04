package main
import (
    "fmt"
    "strings"
)

func main() {
    var s string
    
    fmt.Printf("Digite uma string: ")
    fmt.Scan(&s)
    
    s = strings.ToUpper(s)
    fmt.Printf("%s", s)
}
