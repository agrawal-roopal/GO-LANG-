package main

import "fmt"

func main(){
	num := []int{1,2,3}
	for i,j := range num{
		fmt.Println(i,j)
	}
	for _,j := range num{
		fmt.Println(j)
	}
	for i := range num{
		fmt.Println(i)
	}
}