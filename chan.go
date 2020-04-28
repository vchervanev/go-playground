package main

import "fmt"

var numbers = make(chan string, 100)

func producer(id, count int) {
	for i := 0; i < count; i++ {
		numbers <- fmt.Sprintf("producer[%d] -- %d", id, i)
	}
}

func consumer(id int) {
	for value := range numbers {
		fmt.Printf("consumer[%d]: %s\n", id, value)
	}
}
func main() {
	go producer(100, 10)
	go producer(200, 10)
	go consumer(999)
	go consumer(888)
	consumer(777)
}
