package main
import "fmt"

type Message struct {
    x string
    y *string
}

func (v Message) Print () {
    if v.y != nil {
        fmt.Println (v.x, *v.y)
    } else {
        fmt.Println (v.x)
    }
}

func (v *Message) Store (x, y string) {
    v.x = x
    v.y = &y
}

func main() {
    m := &Message {}
    m.Print ()
    m.Store("Hello", "world")
    m.Print ()
}


