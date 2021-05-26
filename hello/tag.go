package main
import (
	"fmt"
	"reflect"
	)
type Animal struct{
	name string	`required max:"100"`
	origin string
}


func main(){
	t :=reflect.TypeOf(Animal{})
	field,_ := t.FieldByName("name")
	fmt.Println(field)
}