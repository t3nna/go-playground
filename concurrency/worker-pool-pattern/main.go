package main

import (
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// read implements the worker pool pattern from your book.
// It is modified to accept the pool size 'n' and the 'taskFunc' to perform.
func read(r io.Reader, taskFunc func([]byte) int, n int) (int64, error) {
	var count int64
	var wg sync.WaitGroup

	// Create a buffered channel with a capacity equal to the pool size.
	// This helps avoid the sender (main goroutine) from blocking
	// if the workers are busy.
	ch := make(chan []byte, n)
	wg.Add(n) // Add 'n' workers to the wait group.

	// --- Start the Worker Pool ---
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done() // Signal that this worker is done when the channel closes.

			// This loop receives tasks from 'ch' until the channel is closed.
			for b := range ch {
				v := taskFunc(b) // Perform the actual task
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	// --- Read Loop (The "Producer") ---
	// This loop reads from the 'io.Reader' and sends tasks to the workers.
	for {
		// Create a new buffer for each read.
		// Note: In a real app, you might use a sync.Pool for high performance
		// to avoid allocating so many buffers.
		b := make([]byte, 1024)

		nRead, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break // End of file, stop sending tasks.
			}
			return 0, err // A real error occurred.
		}

		// Send the valid data (only the part we read) to the channel.
		// The book's example sends the full 1024 bytes, we'll send the valid slice.
		ch <- b[:nRead]
	}

	// --- Shutdown ---
	close(ch) // Close the channel; this signals workers to stop (for range ends).
	wg.Wait() // Wait for all 'n' workers to call wg.Done().

	return count, nil
}

// --- Task Functions for Testing ---

// taskCPU simulates a CPU-BOUND workload (e.g., hashing, calculations).
// It returns 1 (to represent 1 task completed).
func taskCPU(b []byte) int {
	// Simulate "heavy" work by looping
	// This is a simple, non-optimizing loop to consume CPU.
	sum := 0
	for i := 0; i < 1_000_000; i++ {
		sum += (i * len(b)) % 255
	}
	return 1 // We're just counting tasks, so return 1.
}

// taskIO simulates an I/O-BOUND workload (e.g., API call, DB query).
// It returns 1 (to represent 1 task completed).
func taskIO(b []byte) int {
	// Simulate waiting for a network/database response.
	time.Sleep(50 * time.Millisecond)
	return 1
}

// --- Main Function to Run the Tests ---

func main() {
	// Create some sample data with 100 "tasks" (separated by '|').
	// A "task" is just some data we read. Our reader will read 1024 bytes
	// at a time, so we need enough data.
	taskData := strings.Repeat("some-data-packet|", 200) // ~4KB of data

	// --- Test 1: CPU-Bound Workload ---
	fmt.Println("--- Starting CPU-Bound Test ---")

	// As the book recommends: pool size = number of logical CPUs.
	cpuPoolSize := runtime.GOMAXPROCS(0)
	fmt.Printf("Workload: CPU-Bound\nPool Size: %d (runtime.GOMAXPROCS)\n", cpuPoolSize)

	cpuReader := strings.NewReader(taskData)
	startCPU := time.Now()

	countCPU, err := read(cpuReader, taskCPU, cpuPoolSize)

	fmt.Printf("CPU-Bound Test complete in: %v\n", time.Since(startCPU))
	fmt.Printf("Total tasks processed: %d\n", countCPU)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n---------------------------------\n")

	// --- Test 2: I/O-Bound Workload ---
	fmt.Println("--- Starting I/O-Bound Test ---")

	// We pick an arbitrary *high* number, much larger than GOMAXPROCS.
	// We can run 50 "simultaneous" I/O calls.
	ioPoolSize := 50
	fmt.Printf("Workload: I/O-Bound (50ms wait per task)\nPool Size: %d\n", ioPoolSize)

	ioReader := strings.NewReader(taskData)
	startIO := time.Now()

	countIO, err := read(ioReader, taskIO, ioPoolSize)

	fmt.Printf("I/O-Bound Test complete in: %v\n", time.Since(startIO))
	fmt.Printf("Total tasks processed: %d\n", countIO)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
