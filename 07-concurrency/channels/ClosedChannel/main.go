package main

import "fmt"

func main() {
	linkChannel := make(chan string)
	go ping(linkChannel)

	for {
		_, ok := <-linkChannel
		if ok {
			fmt.Println("Channel Open")
		} else {
			fmt.Println("Channels closed")
			break
		}
	}
}

func ping(c chan string) {
	links := []string{
		"https://www.golinuxcloud.com/", "https://www.tesla.com/", "https://www.google.com/",
	}

	for _, links := range links {
		c <- links
	}

	close(c)
}
