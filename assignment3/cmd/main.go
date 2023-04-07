package main

import (
  www"assignment3/pkg/web"
  "fmt"
)

func main() {
  fmt.Print("Server is running on localhost 8181")
  www.Connect()
} 
