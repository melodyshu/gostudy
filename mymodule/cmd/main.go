package main

import (
    "fmt"
    "github.com/yourusername/mymodule/greetings"
)

func main() {
    message := greetings.Hello("World")
    fmt.Println(message)
}
