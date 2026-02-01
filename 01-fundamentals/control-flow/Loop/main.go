package main

import "fmt"

func main() {

	a := 6
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	if a > 3 && a < 6 {
		fmt.Println("a is within range")
	}

	c := 1
	b := 2

	if c > b {
		fmt.Println("c is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}

	// For loop

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break
		}
	}

	// continue
	for i := i; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Nested for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("Outer loop iteration %d\n", i)
		for j := 0; j < 2; j++ {
			fmt.Printf("i = %d j = %d \n", i, j)
		}
	}

	// There is not while loop .... same can be achieved using for loop

	k := 1
	for k <= 5 {
		fmt.Println(k)
		k++
	}

	// For range loop

	letters := []string{"a", "b", "c", "d", "e"}

	for i, x := range letters {
		fmt.Printf("The values are %d and %s \n", i, x)
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	//Only index
	fmt.Println("\nOnly Index")
	for i := range letters {
		fmt.Printf("Index: %d\n", i)
	}

	//Without index and value. Just print array values
	fmt.Println("\nWithout Index and Value")
	x := 0
	for range letters {
		fmt.Printf("Index: %d Value: %s\n", i, letters[x])
		x++
	}

	// Channel Iteration
	fmt.Println("Channel Iterations")
	ch := make(chan int)

	go func(){
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for n := range ch{
		fmt.Println(n)
	}

	// Map iterations
	fmt.Println("Map Iterations")
	m := map[string]int{
		"one" : 1,
		"two" : 2,
		"three" : 3, 
	}

	for k, v := range m{
		fmt.Println(k ,v)
	}
}
