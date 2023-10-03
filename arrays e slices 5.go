package main
import "fmt"

func main() {
    var i, j, n int
    var m [2][3]int
    
    for i = 0; i < 2; i++ {
        for j = 0; j < 3; j++ {
            fmt.Printf("[%d][%d]: ", i, j)
            fmt.Scan(&n)
            
            m[i][j] = n
        }
    }
    
    for i = 0; i < 2; i++ {
        for j = 0; j < 3; j++ {
            fmt.Printf("%d  ", m[i][j])
        }
        fmt.Println()
    }
}
