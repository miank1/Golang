package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

}
