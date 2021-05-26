package main
import "fmt"

func main(){
	var input int
	LOOP:
	fmt.Println("enter even no.")
	fmt.Scanln(&input)
	if input%2 != 0{
		goto LOOP
	}else{
		fmt.Println("even")
	}
}