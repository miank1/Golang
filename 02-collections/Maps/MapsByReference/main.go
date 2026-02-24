package main

import "fmt"

func update(m map[string]int) {
	m["newKey"] = 42
}

func modifyMap(m map[string]int) {
	m["Key1"] = 100
	m["Key2"] = 100
}

func main() {

	m1 := make(map[string]int)

	m1["Key1"] = 10
	m1["Key2"] = 20

	fmt.Println("Original Map")
	modifyMap(m1)

	fmt.Println("Map after modification")
	fmt.Println(m1)

	m := map[string]int{"a": 1}
	update(m)
	fmt.Println(m)

}
