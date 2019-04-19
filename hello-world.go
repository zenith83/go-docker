package main

import(
  "fmt"
  "runtime"
)

func main(){
  fmt.Printf("hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}