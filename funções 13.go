package main
import (
    "fmt" 
)

func chamada() {
    n := 0
    soma := 0

    for {
        fmt.Scanln(&n)
        
        for n > 0 {
            soma += n % 10
            n /= 10
        }
        
        fmt.Println(soma)
        n = 0
        soma = 0
    }
}

func main() {
    chamada()
}
