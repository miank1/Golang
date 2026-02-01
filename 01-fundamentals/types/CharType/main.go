package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var c byte = 'a'
	fmt.Println("Printing the byte")
	fmt.Printf("Size : %d\n Type: %s\nCharacter: %c", unsafe.Sizeof(c), reflect.TypeOf(c), c)

	r := '£'
	fmt.Println("\nPrinting Rune:")
	//Print Size, Type, CodePoint and Character
	fmt.Printf("Size: %d\nType: %s\nUnicode CodePoint: %U\nCharacter: %c\n", unsafe.Sizeof(r), reflect.TypeOf(r), r, r)

	s := "µ" //Micro sign
	fmt.Println("\nPriting String:")
	fmt.Printf("Size: %d\nType: %s\nCharacter: %s\n", unsafe.Sizeof(s), reflect.TypeOf(s), s)

}
