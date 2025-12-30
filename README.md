# Go Advanced Concepts Playground

A collection of Go examples and experiments exploring advanced concepts and patterns.

## Structure

### Concurrency
- **goroutines-channels**: Basic goroutines and channel patterns
- **nil-channel**: Working with nil channels
- **select**: Using select statement for channel operations
- **stop-go-routine**: Patterns for gracefully stopping goroutines
- **syncCond**: Using sync.Cond for condition variables
- **worker-pool-pattern**: Worker pool implementation

### Context
Examples of using the `context` package for cancellation and timeouts.

### Errors
Error handling patterns and best practices.

### Loops
- **rangeLoop**: Range loop examples
- **rangeLoopPointers**: Range loop with pointers

### Fundamentals
- **map**: Map operations and patterns
- **receivers**: Method receivers (value vs pointer)
- **slices**: Slice operations and internals
- **strings**: String manipulation examples

### Kata
Solutions to CodeWars problems.

## Usage

Each directory contains a `main.go` file that can be run independently:

```bash
go run <directory>/main.go
```

## Module

```
module advanced-concepts
```

