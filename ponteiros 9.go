package main
import (
    "fmt"
    )

type Livro struct {
    titulo string
    autor string
    preco float32
}

func ponteiro(p *Livro) {
    p.preco *= 0.9
}

func main() {
    var u Livro
    
    u.titulo = ""
    u.autor = "Rodrigo"
    u.preco = 111
    
    fmt.Println(u)
    ponteiro(&u)
    fmt.Println(u)
}
