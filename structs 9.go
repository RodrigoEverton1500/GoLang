package main
import (
    "fmt"
    )

type Aluno struct {
    nome string
    idade int
    notas [] float32
}

func menu(class Aluno) {
    var op int
    
    fmt.Printf("1- Modificar\n2- Média\n3- Imprimir\n")
    fmt.Print("Operação: ")
    fmt.Scan(&op)
    
    switch op {
    case 1:
        modificar(class)
    case 2:
        media(class)
    case 3:
        imprimir(class)
    default:
        main()
    }
}

func modificar(class Aluno) {
    var n int
    var op int
    
    fmt.Printf("1- Adicionar\n2- Remover \n")
    fmt.Print("Operação: ")
    fmt.Scan(&op)
    
    switch op {
    case 1:
        fmt.Print("Nota: ")
        fmt.Scan(&n)
        class.notas = append(class.notas, float32(n))
    case 2:
        fmt.Print("Indice da nota: ")
        fmt.Scan(&n)
        class.notas = append(class.notas[:n-1], class.notas[n:]...)
    }
    
    menu(class)
}

func media(class Aluno) {
    l := len(class.notas)
    var media float32
    
    for _, i := range class.notas {
        media += i
    }
    media /= float32(l)
    
    fmt.Printf("Média: %.2f\n", media)
    menu(class)
}

func imprimir(class Aluno) {
    fmt.Println(class)
    menu(class)
}

func main() {
    var a Aluno
    var n float32
    s := [] float32 {}
    
    fmt.Print("Nome: ")
    fmt.Scan(&a.nome)
    fmt.Print("Idade: ")
    fmt.Scan(&a.idade)
    for {
        fmt.Print("Nota: ")
        fmt.Scan(&n)
        if n == -1 {
            break
        }
        s = append(s, n)
    }
    a.notas = s
    
    menu(a)
}
