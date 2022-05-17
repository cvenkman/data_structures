package main

import (
	"fmt"
	"log"

	DS "github.com/cvenkman/data_structures"
)

func main() {
    q := DS.Queue{}
    q.Enqueue("cat")
    fmt.Println("stack after enqueue cat:", (q)) // {[cat]}
    fmt.Println("empty stack:", q.Empty()) // false
    value, err := q.Dequeue()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("dequeue value:", value) // 4
    fmt.Println("stack after dequeue:", q) // {[]}
    
    if _, err = q.Dequeue(); err != nil {
        log.Fatal(err) // empty stack
    }
}