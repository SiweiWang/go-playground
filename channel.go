// The challenge with concurrent programming stems from sharing data.
// If goroutines share no data, it does not need synchronizing them.

// Channels help male concurrent programming saner by taking shared data out of the picture
// A channel is a commnuication pipe between goroutines which is used to pass data.
// Since one goroutine can pass data to another goroutine via channel, at any point in time, only one goroutine has access to the data.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer fmt.Println("existing ...")
	// create a channel
	c := make(chan int)

	//
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)

	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 1000)
	}
}
