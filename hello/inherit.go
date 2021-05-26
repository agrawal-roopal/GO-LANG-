package main
import "fmt"

type Animal struct{
	name string
	origin string
}
type Bird struct{
	Animal
	speed float32
	canFlay bool
}

func main(){
	b :=Bird{}
	b.name = "emu"
	b.origin = "aust"
	b.speed = 48
	b.canFlay =false
	fmt.Println(b)
}