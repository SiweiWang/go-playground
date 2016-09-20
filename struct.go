package main

import (
  "fmt"
)

type Saiyan struct {
  Name string
  Power int
}
func main() {
  type Saiyan struct {
    Name string,
    Power int,
  }

  fmt.Println(Saiyan)
  fmt.Println(&Saiyan)
}