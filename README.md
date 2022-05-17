## Data Structures (containers)
    - [Stack](#Stack)
    - [Queue](#Queue)

import for all main functions:
``` go
import (
    "fmt"
    "log"
    DS "github.com/cvenkman/data_structures"
)
```

### Stack
Стек — это коллекция, доступ в которой предоставлен только к последнему добавленному элементу.

Пример работы со стеком
``` go
func main() {
    s := DS.Stack{}
    fmt.Println("empty stack:", s.Empty()) // true
    s.Push(3)
    s.Push(4)
    // stack.Push("error") // panic
    fmt.Println("stack after push 3 and 4:", (s)) // {[3 4]}
    value, err := s.Pop()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("pop value:", value) // 4
    fmt.Println("stack after pop:", s) // {[3]}
}
```

### Queue
Очередь — это коллекция, в которой элементы кладутся (enqueue) и забираются (dequeue) с разных концов.
``` go
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
```
