// Slice basics

package main

import "fmt"

func main() {
	s := make([]int, 0, 4)

	fmt.Println("Length ", len(s))
	fmt.Println("Capacity ", cap(s))

	// Appending Slice

	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)

	// Capacity exceeded - new array allocated and old data copied - slice now points to new array.

	s = append(s, 5)
	s = append(s, 6)

	fmt.Println("Length ", len(s))
	fmt.Println("Capacity ", cap(s))

	// Indexing and Slicing

	a := []int{10, 20, 30, 40, 50}
	fmt.Println("a", a)

	b := a[1:4]
	fmt.Println("b", b)

	c := a[:3]
	fmt.Println("c", c)

	d := a[2:]
	fmt.Println("d", d)

	// declare a slice
	a1 := [5]int{1, 2, 3, 4, 5}
	var b1 []int = a1[1:4] // Creating a slice from an array
	fmt.Println(b1)

	c1 := []int{6, 7, 8}
	fmt.Println(c1)

	// changes done in slice is actually done in an array
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before ", darr)
	for i := range dslice {
		dslice[i]++
	}

	fmt.Println("array after", darr)
	numa := [3]int{78, 79, 80}

	nums1 := numa[:] // contains all the elements of the slice
	nums2 := numa[:]
	fmt.Println("Array before change 1 ", nums1)
	fmt.Println("Array before change 1 ", nums2)
	nums1[0] = 100
	fmt.Println("Array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("Array after modification to slice nums2", numa)

	// Length and Capacity of slice
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Printf("length of slice is %d and capacity is %d \n", len(fruitSlice), cap(fruitSlice))
	fruitSlice = fruitSlice[:cap(fruitSlice)]
	fmt.Println("After re-slicing length is ", len(fruitSlice),
		"and capacity is ", cap(fruitSlice))

}
