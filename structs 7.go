package main
import "fmt"

type Animal struct {
    nome string
    especie string
    idade int
    som string
}

func som(class Animal) Animal {
    var s string
    
    fmt.Print("Novo som: ")
    fmt.Scan(&s)
    
    if s == "0" {
        return class
    } else {
        class.som = s
        return class
    }
}

func imprimir(class Animal) {
    fmt.Print(class)
}

func main() {
    var u Animal
    
    fmt.Print("Nome:")
    fmt.Scan(&u.nome)
    fmt.Print("Espécie:")
    fmt.Scan(&u.especie)
    fmt.Print("Idade:")
    fmt.Scan(&u.idade)
    fmt.Print("Som: ")
    fmt.Scan(&u.som)
    
    u = som(u)
    imprimir(u)
}
