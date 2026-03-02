package main

func merge(ch1, ch2 <-chan int) <-chan int {

	result := make(chan int, 1)

	go func() {
		defer close(result)

		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
				} else {
					result <- v
				}
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
				} else {
					result <- v
				}
			}
		}

	}()

	return result

}

func main() {

}
