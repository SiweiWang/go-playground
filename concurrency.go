// Two powerful mechanisms: goroutunes and channels

//Goroutines
/* A goroutine is similar to a thread, but it is schediled by Go, not the OS. 
 * Code that runs in a goroutine can run concurrently with other code.
 */

package main 

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("start")

  // Start a goroutine using the key word go and the function name
  go process()

  // Need a sleep here becuase the main process exists before the goroutine get a chance to execute.
  // To Solve this , we need to coordinate our code. That is where mutex and channel comes in
  time.Sleep(time.Millisecond*10) // This is bad concurrency contorl
  fmt.Println("done")
}

func process() {
  fmt.Println("processing")
}

