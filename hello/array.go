package main
import "fmt"

func main(){
	grades := [3]int{1,2,3}	
	grades1 := [...]int{4,5,6,7}
	fmt.Println(grades)
	fmt.Println(grades1)

	var students [3]string
	fmt.Println(students)
	students[1]="roopal"
	fmt.Println(students)
	fmt.Println(len(students))
	fmt.Println(cap(students))
	// var matrix [3][2]int = [3][2]int{[2]int{1,1},[2]int{0,0},[2]int{0,1}}
	var matrix [3][2]int
	matrix[0]=[2]int{1,1}
	matrix[1]=[2]int{0,0}
	matrix[2]=[2]int{0,1}

	fmt.Println(matrix)

	a := [...]int{1,2,3}
	b := &a
	b[1]=5
	fmt.Println(a)
	fmt.Println(b)



	//SLICE

	c := []int{3,4,6,7,8,8,9,0}
	fmt.Println(len(c))
	fmt.Println(cap(c))
	v := c[:]	//slice all elements
	h := c[3:]	//slice from 4th elemnet
	i := c[:6]	//slice first 6 elelment
	j := c[3:6]	//slice 4th 5th 6th element
	fmt.Println(v)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

}