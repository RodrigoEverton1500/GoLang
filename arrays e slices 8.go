package main
import "fmt"

func main() {
    var s string
    var i int
    lista := [] string {"a","b","c","d","e","f","g","h"}
    
    fmt.Printf("Digite um valor string: ")
    fmt.Scan(&s)
    
    for i = 0; i < 8; i++ {
        if lista[i] == s {
            lista = append(lista[:i], lista[i+1:]...)
            break
        }
    }
    
    fmt.Printf("%s", lista)
}
