// Another option is to timeout, We are willing for block for sometime, but not forever.
// This is also something easy to achieve 

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
    select {
      case c <- rand.Int():
      // optinal code for normal case
      case t := <- time.After(time.Millisecond * 100):
        fmt.Println("timed out at", t)
    }
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

// time.after returns a channel, so we can select from it.
// The channel is written to after the time expires.

// Rule for select channel
// 1. The first available channel is chosen.
// 2. If multiple channels are available, one is randomly picked.
// 3. If no channel is available, the default case is executed
// 4. If there's no default, select blocks