package main

import "fmt"

//function are mainly 4 type
//1. parameter with return type'
//2. parameter with  no return types
//3. no parameter with  retun types
//4. no parametr with no return types

func main() {
	fmt.Println(myfriend(10))
	myfriend2(3)
	fmt.Println(myfrind3())
	myfriend4()
	myfunc(10, 20)
	fmt.Print(myfunc1(20, 30))
}

//1 . Parameter with return types
func myfriend(x int) int {
	return x

}

// 2. Parameter with No retunr types
func myfriend2(x int) {
	fmt.Println(x)
}

//3. No parameter with return types
func myfrind3() int {
	x := 10
	return x

}

// 4. no parameter with no return types
func myfriend4() {
	fmt.Println("Hello")
}

//This is diffrenet Topic for functions
// * functions continuted  = it menas witch we can use shortended parameters
func myfunc(x, y int) { //shortended parameter
	fmt.Println(x, y)
}

// * Multipl result return a functon
func myfunc1(x, y int) (int, int) {
	return x, y
}
