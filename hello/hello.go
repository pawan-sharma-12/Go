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
    names := []string{"Gladys", "Samantha", "Darrin"}
    messages, err := greetings.Hellos(names);
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
}