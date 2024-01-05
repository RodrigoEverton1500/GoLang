package main
import (
    "fmt"
    )

type Livro struct {
    titulo string
    autor string
}

func ponteiro(p *Livro) {
    if p.autor == "anonimo" {
        p.titulo = "desconhecido"
    }
}

func main() {
    var u Livro
    
    u.titulo = ""
    u.autor = "anonimo"
    
    fmt.Println(u)
    ponteiro(&u)
    fmt.Println(u)
}
