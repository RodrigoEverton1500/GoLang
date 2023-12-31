package main
import (
    "fmt"
    )

type Pessoa struct {
    nome string
    idade int
    Endereço
}

type Endereço struct {
    numero int
    rua string
    cidade string
    estado string
}

func f(class Pessoa) {
    n := class.Endereço
    fmt.Println(n)
}

func main() {
    var u Pessoa
    
    u.nome = "Rodrigo"
    u.idade = 18
    u.numero = 000
    u.rua = "100"
    u.cidade = "Brasília"
    u.estado = "DF"
    
    f(u)
}
