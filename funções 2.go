package main
import ("fmt"
        "strings")

func vogal(s string) {
    a := strings.Count(s, "a")
    e := strings.Count(s, "e")
    i := strings.Count(s, "i")
    o := strings.Count(s, "o")
    u := strings.Count(s, "u")
    
    vogal := a+e+i+o+u
    fmt.Printf("Número de vogais: %d", vogal)
}

func main() {
    var s string
    
    fmt.Printf("Digite uma string: ")
    fmt.Scan(&s)
    
    vogal(s)
}
