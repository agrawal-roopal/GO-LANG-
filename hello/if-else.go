package main
import "fmt"

func main(){
	fmt.Print("enter number..")
	var input int
	fmt.Scanln(&input)
	fmt.Print(input)
	fmt.Print("\n")
	if(input % 2 == 0){
		fmt.Println("even \n")
	}else{
		fmt.Print("odd")
	}
}