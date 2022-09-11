package main

import (
	"fmt"
)

func say2(text string, c chan string) {
	c <- text
}

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 1)
	c2 := make(chan string, 2)

	c2 <- "Que"
	c2 <- "Onda"

	fmt.Println(len(c2), cap(c2))

	fmt.Println("Hello")

	go say2("World", c)
	fmt.Println(<-c)

	// Close y Range
	close(c2)
	for message := range c2 {
		fmt.Println(message)
	}

	//select
	email := make(chan string)
	email2 := make(chan string)

	go message("Mensaje1", email)
	go message("Mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email de email", m1)
		case m2 := <-email2:
			fmt.Println("Email de email2", m2)
		}
	}
}
