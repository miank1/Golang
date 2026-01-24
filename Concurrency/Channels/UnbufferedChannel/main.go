package main

func main() {

	ch := make(chan int)

	ch <- 1 // blocks forever

}
