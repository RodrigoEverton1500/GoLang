package main
import "fmt"

func con(slice[]string, c int) {
    var s string
    
    for i := 0; i < c; i++ {
        s += slice[i]
    }
    
    fmt.Printf("%s", s)
}

func main() {
    var s string
    var c int
    slice := [] string {}
    
    for {
        fmt.Printf("Digite uma string: ")
        fmt.Scan(&s)
        
        if s == "s" {
            break
        }
        
        c++
        slice = append(slice, s)
    }
    
    con(slice, c)
}
