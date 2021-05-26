package main
import "fmt"


func main(){
	mapEx := make(map[string]int)
	mapEx =map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	delete(mapEx,"a")
	fmt.Println(mapEx["d"])
}