```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var counter int
        const numRoutines = 1000
        var mu sync.Mutex // Add mutex

        for i := 0; i < numRoutines; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Lock the mutex before accessing counter
                        counter++
                        mu.Unlock() // Unlock the mutex after accessing counter
                }()
        }

        wg.Wait()
        fmt.Println("Counter value:", counter)
}
```