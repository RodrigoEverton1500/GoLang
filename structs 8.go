package main
import (
    "fmt"
    )

type Viagem struct {
    origem string
    destino string
    data int
    preco float32
}

func caro(s []Viagem) {
    var t float32
    l := len(s)
    preco := [] float32 {}
    
    for i := 0; i < l; i++ {
        v := s[i]
        preco = append(preco, v.preco)
    }
    
    for i := 1; i < l; i++ {
        for j := 0; j < l-1; j++ {
            if preco[i] < preco[j] {
                t = preco[i]
                preco[i] = preco[j]
                preco[j] = t
            }
        }
    }
    
    fmt.Print(preco[l-1])
}

func main() {
    var v1 Viagem
    var v2 Viagem
    var v3 Viagem
    
    v1.origem = "Brasília"
    v1.destino = "São Paulo"
    v1.data = 02012024
    v1.preco = 1000
    v2.origem = "Brasília"
    v2.destino = "São Paulo"
    v2.data = 02012024
    v2.preco = 800
    v3.origem = "Brasília"
    v3.destino = "São Paulo"
    v3.data = 02012024
    v3.preco = 900
    
    s := [] Viagem {v1, v2, v3}
    
    caro(s)
}
