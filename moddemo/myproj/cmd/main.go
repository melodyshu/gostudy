package main

import (
    "fmt"
    "myproj/greetings"
)

func main() {
    message := greetings.Hello("World")
    fmt.Println(message)
}
