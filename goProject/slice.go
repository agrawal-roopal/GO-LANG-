package main
import "fmt"

func main(){
	//slice to group values of same type
	// x := type{values} //composite literals
	x := []int{1,2,3,49}
	fmt.Println(x[2])
	fmt.Println(len(x))
	for i,j := range x {
		fmt.Println(i,j)
	}
	fmt.Println(x[2:4])
	x = append(x,11,22,33)
	fmt.Println(x)
	y := []int{33333343,2455,5454}
	x = append(x, y...)
	fmt.Println(x)

	//delete
	x = append(x[:1], x[3:]...)
	fmt.Println(x)

	z := make([]int, 10, 11)
	fmt.Println(z)
	fmt.Println(len(z))//10
	fmt.Println(cap(z))//11

	z =append(z,876578)
	fmt.Println(z)
	fmt.Println(len(z)) //11
	fmt.Println(cap(z)) //11

	z =append(z,789087)
	fmt.Println(z)
	fmt.Println(len(z))//12
	fmt.Println(cap(z)) //22
	//multi-dimentional 
	xx := []string{"a","b","c"}
	yy := []string{"aa","bb","cc"}
	zz := [][]string{xx,yy}
	fmt.Println(zz)

	for i,j:= range zz{
		fmt.Println("slice :",i)
		for k,l:=range j{
			fmt.Printf("\t index: %v \t value: %v \n",k,l)
		}
	}
	slice_ex()

}

func slice_ex(){
	x := map[string][]string{
		"color": []string{"red","green","blue"},
		"lang": []string{"eng","hindi","french"},
	}
	for i,j := range x{
		fmt.Println("type : ",i)
		for _, l:= range j{
			fmt.Println(l)
		}
	}
}