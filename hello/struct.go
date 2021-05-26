package main
import "fmt"

type Doctor struct{
	id int
	name string
	className []string
} 

func main(){
	aUser := struct{name string}{name:"hiii"}
	aDoctor := Doctor{
		1,
		"harry",
		 []string{
			"hi",
			"there",
		},
	}
	fmt.Println(aDoctor);
	fmt.Println(aUser);
}