package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) // Report the err and stop it
	}
	fmt.Println(input)
}
