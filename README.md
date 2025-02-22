# Go Race Condition Example
This repository demonstrates a simple race condition in Go that occurs when multiple goroutines increment a shared counter without proper synchronization. The counter's final value is likely to be less than expected due to data races.

## How to reproduce
1. Clone this repository
2. Navigate to the project directory
3. Run `go run bug.go`

The output will show a counter value that's likely lower than the expected value (1000 in this case).

## Solution
The solution involves using a `sync.Mutex` to protect the counter from concurrent access.  See `bugSolution.go` for the fix.