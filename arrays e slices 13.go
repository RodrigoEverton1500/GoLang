package main
import "fmt"

func main() {
    var n int
    array := [7] int {1,2,3,4,5,6,7}
    
    fmt.Printf("Digite um valor para o primeiro elemento: ")
    fmt.Scan(&n)
    array[0] = n
    
    fmt.Printf("Digite um valor para último elemento: ")
    fmt.Scan(&n)
    array[6] = n
    
    fmt.Printf("%d", array)
}
