package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("The array is ", arr)

	sample := [2]int{1, 2}
	fmt.Printf("Sample1 : Len %d %v \n", len(sample), sample)

	sample2 := [...]int{1, 2}
	fmt.Printf("Sample2 : Len %d %v \n", len(sample2), sample2)

	sample3 := [2]int{}
	fmt.Printf("Sample3: Len: %d %v \n", len(sample3), sample3)

	sample4 := [...]int{}
	fmt.Printf("Sample3: Len: %d %v \n", len(sample4), sample4)

	// Array of string
	var arr1 [3]string
	fmt.Println(arr1)

	arr1[0] = "abc"
	arr1[1] = "def"
	arr1[2] = "ghi"

	for _, c := range arr1 {
		fmt.Println(c)
	}

	// Array is actually a copy of the existing array
	arr2 := [...]int{1, 2, 3}
	arr3 := &arr2
	arr3[1] = 5
	fmt.Println(arr2)
	fmt.Println(arr3)

}
