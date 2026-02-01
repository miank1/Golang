package main

import (
	"fmt"
	"path"
)

func main() {

	dir, file := path.Split("css/main.css")

	fmt.Println("Dir:", dir)
	fmt.Println("file:", file)
}
