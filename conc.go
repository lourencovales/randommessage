package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

var message = []string{
	"This is message number one",
	"And this is message number two",
	"Let's go for number three",
	"What about number four?",
	"It's a-me, five!",
	"Six? That's two-times-three",
	"Seven is prime",
	"But eight isn't",
	"Nine for sure can be devided by 3",
	"And then there's ten!",
}

var n int

func main() {
	userInput()
	msgBuff := messageOutput(n)

	var wait sync.WaitGroup
	wait.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wait.Done()
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
			fmt.Println(<-msgBuff)
		}()
	}
	wait.Wait()
}

func userInput() {
	fmt.Print("How many messages do you want? Pick a number between 1 and 10: ")
	fmt.Scan(&n)

	if n > 10 || n < 1 {
		err := fmt.Errorf("Number is not between 1 and 10, start again.")
		panic(err)
	}
}

func messageOutput(n int) <-chan string {
	// buffered channel that will keep the messages in random order
	channel := make(chan string, n)
	for i := 0; i < n; i++ {
		a := rand.Intn(len(message))
		channel <- message[a]
		// to guarantee uniqueness, we need to delete the msg that was used
		message = slices.Delete(message, a, a+1)
	}
	close(channel)
	return channel

}
