package main
import "fmt"

func main(){
	a := 10
	b := 3
	fmt.Println(a + b) //13
	fmt.Println(a - b)  //7
	fmt.Println(a / b) //3
	fmt.Println(a * b)	//30
	fmt.Println(a % b)	//1
	fmt.Println(a & b)	//2
	fmt.Println(a | b)	//11
	fmt.Println(a ^ b)	//9
	fmt.Println(a &^ b)	//8

	c := 8 //2^3
	fmt.Println(c << 3)	// 2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3)	//2^3 / 2^3 = 2^0 = 1

	//floating point num
	//(+, -, *, /)
	d := 3.14
	d = 3.17e72 //3.17 * 10^72
	d = 3.17E72 //3.17 * 10^72
	fmt.Printf("%v - %T \n",d,d)
	//complex
	//(+, -, *, /)
	var n complex64 = 1 + 2i
	var m complex64 = complex(2,3)
	fmt.Printf("%v , %T \n",n,n)
	fmt.Printf("%v , %T \n",real(n),real(n)) //1 , float32
	fmt.Printf("%v , %T \n",imag(n),imag(n)) //2 , float32

	fmt.Printf("%v , %T \n",m,m)
	fmt.Printf("%v , %T \n",real(m),real(m)) //1 , float32
	fmt.Printf("%v , %T \n",imag(m),imag(m)) //2 , float32

	//string
	s :="hi there"
	t :="good luck"
	fmt.Printf("%v , %T \n",s+t,s+t) //concat

	z := []byte(s)
	fmt.Printf("%v , %T \n",z,z) //byte value for each char

	//rune
	var r rune = 'a' // r := 'a'
	fmt.Printf("%v , %T \n",r,r) //97 , int32

}