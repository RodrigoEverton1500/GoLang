package main
import (
    "fmt"
    )

type Conta struct {
    titular string
    saldo float32
}

func ponteiro(p *Conta) {
    fmt.Print("Saldo: ")
    fmt.Scan(&p.saldo)
    fmt.Println(*p)
}

func main() {
    var u Conta
    
    u.titular = "Rodrigo"
    u.saldo = 0
    
    ponteiro(&u)
}
