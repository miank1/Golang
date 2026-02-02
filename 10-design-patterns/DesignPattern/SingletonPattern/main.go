package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var once sync.Once

type single struct {
}

var (
	singleInstance,
	singleInstance_1,
	singleInstance_2 *single
)

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created 1.")
		}
	} else {
		fmt.Println("Single instance already created 2.")
	}

	return singleInstance
}

func getInstance_1() *single {
	if singleInstance_1 == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now with Once")
				singleInstance_1 = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance_1
}

func getInstance_2() *single {
	once.Do(func() {
		fmt.Println("Creating Singleton instance...")
		singleInstance_2 = &single{}
	})
	return singleInstance_2
}

func main() {
	for i := 0; i < 5; i++ {
		go getInstance()
	}

	for i := 0; i < 5; i++ {
		go getInstance_1()
	}

	s1 := getInstance_2()
	fmt.Println(s1)

	s2 := getInstance_2()
	fmt.Println(s2)

	// Confirm that both instances are the same
	fmt.Println("Are instances equal?", s1 == s2) // Output: true
	fmt.Scanln()
}
