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
