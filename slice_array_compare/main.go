package main

import "fmt"

func main() {
	// array can be compared to array
	a1 := [3]string{"a", "b", "c"}
	a2 := [3]string{"c", "d", "e"}
	a3 := [3]string{"a", "b", "c"}
	fmt.Println(a1 == a2, a1 == a3)

	// slice can only be compared to nil
	// a1 := []string{"a", "b", "c"}
	// a2 := []string{"c", "d", "e"}
	// a3 := []string{"a", "b", "c"}
	// fmt.Println(a1 == a2, a1 == a3)

	//"Only" a runtime panic: runtime error: index out of range
	// s := make([]int,3)
	// s[3] = 3

	//Compile-time error:invalid array index 3(out of bounds for 3-element array)
	// a := [3] int{}
	// a[3] = 3
	var a = []int{0, 1, 2, 3, 4, 5}
	k := a[3]
	fmt.Print(k)
}
