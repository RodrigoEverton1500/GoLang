package main
import "fmt"

func main() {
    var peso float32
    
    fmt.Printf("Digite peso em Kg: ")
    fmt.Scan(&peso)
    
    fmt.Printf("Peso em libras: %.2f", peso*2.205)
}
