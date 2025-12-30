package main

import (
	"fmt"
	"sync"
	"time"
)

// watcher holds resources that need graceful cleanup.
type watcher struct {
	wg   sync.WaitGroup
	quit chan struct{} // A channel to signal the goroutine to stop
}

// newWatcher creates a watcher, starts its goroutine, and returns it.
func newWatcher() *watcher {
	w := &watcher{
		// Make the 'quit' channel
		quit: make(chan struct{}),
	}

	// Add 1 to the WaitGroup *before* starting the goroutine.
	// This prevents a race condition where close() might run before Add().
	w.wg.Add(1)

	// Start the background goroutine
	go w.watch()

	fmt.Println("Main: newWatcher() created and goroutine started.")
	return w
}

// watch is the background goroutine's main loop.
// It listens for work or a quit signal.
func (w *watcher) watch() {
	// Defer wg.Done() to signal that this goroutine has
	// *officially* finished when the function returns.
	defer w.wg.Done()

	fmt.Println("   [Goroutine]: watch() started.")

	// Simulate doing work every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop() // Clean up the ticker's resources

	for {
		select {
		case <-ticker.C:
			// This is the "work"
			fmt.Println("   [Goroutine]: ...doing work (watching)...")

		case <-w.quit:
			// We received a stop signal from the close() method.
			fmt.Println("   [Goroutine]: Quit signal received.")

			// --- This is the critical cleanup phase ---
			fmt.Println("   [Goroutine]: Cleaning up resources (e.g., closing DB conn)...")
			time.Sleep(250 * time.Millisecond) // Simulate time taken for cleanup
			fmt.Println("   [Goroutine]: Cleanup complete. Exiting.")

			// Return from the function, which will trigger the 'defer wg.Done()'
			return
		}
	}
}

// close is the public method to shut down the watcher.
// It will block until the 'watch' goroutine is fully stopped.
func (w *watcher) close() {
	fmt.Println("Main: Calling w.close()...")

	// Signal the goroutine to stop by closing the quit channel.
	// All receivers on this channel will get a "zero" value.
	close(w.quit)

	// --- This is the key to the pattern ---
	// Wait for the goroutine to call 'wg.Done()'.
	// This *blocks* 'close()' (and thus 'main') until cleanup is done.
	fmt.Println("Main: Waiting for goroutine to finish...")
	w.wg.Wait()

	fmt.Println("Main: Watcher fully closed.")
}

// --- Main Application ---

func main() {
	fmt.Println("Main: Application starting...")

	// Create the watcher (which starts the goroutine)
	w := newWatcher()

	// This is the solution! We use 'defer' to call 'close()'.
	// This guarantees that close() is called before main exits.
	// Because w.close() *blocks*, 'main' cannot exit until
	// the goroutine is confirmed to be finished.
	defer w.close()

	// Simulate the main application running for a short time
	fmt.Println("Main: Application running for 2 seconds...")
	time.Sleep(2 * time.Second)

	fmt.Println("Main: Application shutting down.")
	// The 'defer w.close()' will execute here
}
