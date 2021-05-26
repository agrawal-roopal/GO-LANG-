package main
import "fmt"

func main(){
	mm :=map[string]int{
		"aa":1,
		"bb":2,
	}
	fmt.Println(mm);
	fmt.Println(mm["aa"]);
	cc, kk :=mm["cc"]
	fmt.Println(cc,kk)  ///0 , false
	if cc, kk :=mm["aa"]; kk{
		fmt.Println(cc)
	}
	//add elemnt
	mm["dd"] = 9
	for x,y := range mm{
		fmt.Println(x,y)
	}

	//delete

	delete(mm,"aa")
	fmt.Println(mm)
}