package main
import "fmt"

func main(){
	k :=20
	switch k {
		case 10:
			fmt.Print("10 \n")
		case 20:
			fmt.Print("20 \n");fallthrough;
		case 30:
			fmt.Print("30 \n");fallthrough;
		default:
			fmt.Print("default \n")
	}
}

//20 
//30 
//default 