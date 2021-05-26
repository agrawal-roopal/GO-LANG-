package main 
import "fmt"

func main(){
	var a int = 42
	// var b *int = &a
	b := &a
	fmt.Printf("%v\n", b)
	a = 12
	fmt.Println(a , *b)
	*b= 13
	fmt.Println(a , *b)

}