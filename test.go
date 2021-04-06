package main

import (
	"fmt"
	"math"
)

var a, b, c = 0, 1, 2
var str1, str2, str3 = "aaa", "bbb", "ccc"

func test(num1, num2 int) (int, int) {
	return num2, num1
}

func test2(input int) (output1, output2 int) {
	output1 = input + 2
	output2 = 5
	return
}

func test3() int {
	value1, value2 := 22, 23

	ans := 5 + value1 + value2

	return ans
}

func test4() {
	var v float32
	v = 1.5 // change me!
	fmt.Printf("v is of type %T\n", v)
}

func myloop(upto int) {
	temp := 0
	for temp < upto {
		temp += 1
		fmt.Println(temp)
	}
}

func myloop2() {
	temp := 0
	for {
		temp++
		fmt.Println(temp)
	}
}

func Sqrt(x float64) (z float64) {

	z = 1.0

	old_z := 999.0

	diff := 999.0

	for diff > 0.0001 {

		old_z = z

		z = z - (math.Pow(z, 2)-x)/(2*z)

		diff = math.Abs(z - old_z)

		fmt.Println("z=", z)
		fmt.Println("diff=", diff)

	}

	return
}

func waiting() {
	defer fmt.Println("waiting......")

	fmt.Println("a")

	fmt.Println("b")

	fmt.Println("c")

}

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(p)  // address of i
	fmt.Println(*p) // value of i

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	fmt.Println(p) // address of j
	*p = *p / 37   // divide j through the pointer

	fmt.Println(j) // see the new value of j
}

type myStruct struct {
	X int
	Y int
	z int
}

func (m myStruct) add() int {
	return m.X + m.Y + m.z
}

func main() {
	test := myStruct{1, 2, 6}

	fmt.Println(test.add())

}
