package main

import (
  "fmt"
  "os"
)

func main() {
  sceneFile := os.Args[len(os.Args)-1]
  fmt.Printf("%s\n", sceneFile)
}