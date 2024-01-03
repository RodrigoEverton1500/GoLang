package main
import (
    "fmt"
    )

type Filme struct {
    titulo string
    diretor string
    ano int
    avaliacoes []float32
}

func menu(f Filme) {
    var op int
    
    fmt.Printf("1- Modificar\n2- Média\n3- Imprimir\n")
    fmt.Print("Operação: ")
    fmt.Scan(&op)
    
    switch op {
    case 1:
        modificar(f)
    case 2:
        media(f)
    case 3:
        imprimir(f)
    default:
        menu(f)
    }
}

func modificar(f Filme) {
    var n float32
    var op int
    s := f.avaliacoes
    
    fmt.Printf("1- Adicionar\n2- Remover\n")
    fmt.Print("Operação: ")
    fmt.Scan(&op)
    
    switch op {
    case 1:
        fmt.Print("Avaliação: ")
        fmt.Scan(&n)
        s = append(s, n)
    case 2:
        fmt.Print("Indice: ")
        fmt.Scan(&n)
        i := int(n)
        s = append(s[:i-1], s[i:]...)
    }
    
    f.avaliacoes = s
    menu(f)
}

func media(f Filme) {
    var media float32
    
    for _, i := range f.avaliacoes {
        media += i
    }
    
    l := len(f.avaliacoes)
    media /= float32(l)
    fmt.Printf("Média: %.1f\n", media)

    menu(f)
}

func imprimir(f Filme) {
    fmt.Println(f)
    menu(f)
}

func main() {
    var f Filme
    
    f.titulo = "Filme"
    f.diretor = "Rodrigo"
    f.ano = 03022024
    f.avaliacoes = [] float32 {}
    
    menu(f)
}
