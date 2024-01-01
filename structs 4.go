package main
import "fmt"

type Playlist struct {
    nome string
    musicas []Musica
}

type Musica struct {
    titulo string
    artista string
    duracao int
}

func f(class Playlist, s []Musica) {
    var soma int
    
    fmt.Println(class.nome)
    
    for i := 0; i < len(s); i++ {
        m := s[i]
        soma += m.duracao
    }
    
    fmt.Println(soma)
    fmt.Println(s[0])
    fmt.Println(s[1])
    fmt.Println(s[2])
}

func main() {
    var u Playlist
    var msc1 Musica
    var msc2 Musica
    var msc3 Musica
    
    msc1.titulo = "Musica 1"
    msc1.artista = "Joao"
    msc1.duracao = 100
    
    msc2.titulo = "Musica 2"
    msc2.artista = "Pedro"
    msc2.duracao = 150
    
    msc3.titulo = "Musica 3"
    msc3.artista = "Lucas"
    msc3.duracao = 200
    
    musicas := [] Musica {msc1, msc2, msc3}
    
    u.nome = "PlayOne"
    u.musicas = musicas
    
    f(u, musicas)
}
