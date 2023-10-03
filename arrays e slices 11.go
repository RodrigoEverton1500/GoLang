package main
import "fmt"

func main() {
    var i, j, c int
    var m[2][3] int
    
    for i = 0; i < 2; i++ {
        for j = 0; j < 3; j++ {
            c++
            m[i][j] = c
        }
    }
    
    fmt.Printf("Digite linha: ")
    fmt.Scan(&i)
    
    fmt.Printf("Digite coluna: ")
    fmt.Scan(&j)
    
    fmt.Printf("Valor em [%d][%d]: %d", i, j, m[i][j])
}
