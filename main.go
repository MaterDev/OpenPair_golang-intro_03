package main

import (
    "fmt"
    "time"
)

// Structs in Go
// The Animal struct now holds additional fields like Name and Age.
type Animal struct {
    Name string
    Age  int
}

// Speak method for Animal struct.
// Dynamically changes behavior based on the Animal's Name.
func (a Animal) Speak() string {
    switch a.Name {
    case "Dog":
        return "Woof!"  // Returns "Woof!" if the animal is a Dog.
    case "Cat":
        return "Meow!"  // Returns "Meow!" if the animal is a Cat.
    default:
        return "Unknown animal"  // Default case for other animals.
    }
}

// Interfaces in Go
// Mover interface is introduced with a single method Move.
type Mover interface {
    Move() string
}

// Implementing the Mover interface in Animal struct.
// This allows instances of Animal to be treated as Movers.
func (a Animal) Move() string {
    return a.Name + " is moving!"
}

// Demonstrating Interface Usage with a Slice of Interfaces
// Here we are demonstrating how you can use interfaces to handle a collection of different types uniformly.
func processAnimals(animals []Mover) {
    for _, animal := range animals {
        // Each animal in the slice is a Mover, so we can call the Move method.
        fmt.Println(animal.Move())
    }
    // This showcases how interfaces allow us to write functions that can operate on a slice of different types,
    // as long as each type implements the interface methods.
}

// Enhanced Concurrency Example
// Demonstrating buffered channels and select statement.
func demonstrateConcurrencyEnhanced() {
    c := make(chan string, 1) // Creating a buffered channel.

    go func() {
        time.Sleep(1 * time.Second)
        c <- "message from goroutine"
    }()

    select {
    case msg := <-c:
        fmt.Println("Received:", msg)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout: no message received")
    }
}

// Enhanced Error Handling
func divideEnhanced(x, y float64) (float64, error) {
    if y == 0.0 {
        return 0, fmt.Errorf("divide error: cannot divide by zero")
    }
    return x / y, nil
}

func main() {
    // Demonstrating Structs and Methods
    dog := Animal{Name: "Dog", Age: 5}
    fmt.Println(dog.Name, "says:", dog.Speak())

    cat := Animal{Name: "Cat", Age: 3}
    fmt.Println(cat.Name, "says:", cat.Speak())

    // Demonstrating Interfaces and Real-world Use Case
    // We create a slice of Mover, which both dog and cat satisfy.
    animals := []Mover{dog, cat}
    processAnimals(animals)

    // Demonstrating Enhanced Concurrency
    fmt.Println("\nDemonstrating Enhanced Concurrency with Goroutines and Channels")
    demonstrateConcurrencyEnhanced()

    // Demonstrating Enhanced Error Handling
    fmt.Println("\nDemonstrating Enhanced Error Handling in Go")
    result, err := divideEnhanced(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result of division:", result)
}
