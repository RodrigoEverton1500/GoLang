package main
import "fmt"

func main() {
    var n int
    
    fmt.Printf("Digite um valor: ")
    fmt.Scan(&n)
    
    fmt.Printf("\nDobro: %d", n*2)
    fmt.Printf("\nTriplo: %d", n*3)
    fmt.Printf("\nQuadruplo: %d", n*4)
}
