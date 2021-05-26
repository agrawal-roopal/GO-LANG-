package main

import "fmt"
 
const(
	a = iota
	b = iota
	c = iota
)
 
// const(
// 	a1 = iota 
// 	b1 
// 	c1 
// )
const(
	a1 = 1 << (10 * iota)
	b1 
	c1 
)
func main(){
	fmt.Printf("%v,%T",a,a) //0, int
	fmt.Printf("%v,%T",b,b)	//1, int
	fmt.Printf("%v,%T",c,c)	//2, int
	fmt.Printf("%v,%T",a1,a1)	//0, int
	fmt.Printf("%v,%T",b1,b1)	//1, int
	fmt.Printf("%v,%T",c1,c1)	//2, int
}