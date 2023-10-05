package main
import "fmt"

func primeiro(slice[]int, n int) (r int){
    if slice[0] == n {
        r = n
        return r
    } else {
        r = -1
        return r
    }
}

func main() {
    var n, retorno int
    slice := [] int {}
    
    for {
        fmt.Printf("Digite um valor: ")
        fmt.Scan(&n)
        
        if n == 0 {
            break
        }
        
        slice = append(slice, n)
    }
    
    fmt.Printf("Digite valor a ser procurado: ")
    fmt.Scan(&n)
    
    retorno = primeiro(slice, n)
    
    fmt.Printf("%d", retorno)
}
