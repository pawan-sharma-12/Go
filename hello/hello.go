package main

import (
    "fmt"
    "example.com/greetings"
    "log"
)

func main() {
    // Set properties of the predefined Logger, including
    log.SetPrefix(("greetings: "))
    log.SetFlags(0);
    // Get a greeting message check for errors if no then print it.
    message, err := greetings.Hello("Siddharth")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}