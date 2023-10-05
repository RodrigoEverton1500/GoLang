package main
import ("fmt"
"errors")

func con(slice[]string, c int) (erro error){
    var s string
    
    if len(slice) == 0 {
        erro = nil
        return erro
    } else {
        for i := 0; i < c; i++ {
            s += slice[i]
            s += ","
        }
    
        fmt.Printf("%s", s)
        erro = errors.New("ok")
        return erro
    }
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
    
    erro := con(slice, c)
    fmt.Println(erro)
}
