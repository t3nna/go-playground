package main

import (
	"fmt"
	"time"
)

// merge multiplexes two channels into one, using the 'nil channel'
// pattern to gracefully handle the closing of its inputs.
func merge(ch1, ch2 <-chan int) <-chan int {
	// We make the output channel (unbuffered is fine)
	ch := make(chan int)

	go func() {
		// This loop continues as long as *at least one*
		// of the input channels is NOT nil.
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					// --- THE KEY TRICK ---
					// ch1 is closed. We set it to nil.
					// On the next loop, the 'case <-ch1'
					// will block forever, effectively disabling it.
					fmt.Println("   [Merge Goroutine]: ch1 closed, setting to nil.")
					ch1 = nil
					break
				}
				ch <- v

			case v, open := <-ch2:
				if !open {
					// --- THE KEY TRICK ---
					// ch2 is closed. We set it to nil.
					fmt.Println("   [Merge Goroutine]: ch2 closed, setting to nil.")
					ch2 = nil
					break
				}
				ch <- v
			}
		}

		// The loop only exits when BOTH ch1 and ch2 are nil.
		// We can now safely close the merged channel.
		fmt.Println("   [Merge Goroutine]: Both channels nil. Closing merged channel.")
		close(ch)
	}()

	return ch
}

// --- Main Function to Test the Merge ---

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start the merge function. It returns the merged channel.
	merged := merge(ch1, ch2)

	// --- Producer Goroutines ---
	// These simulate two processes sending data at different rates.

	// Producer 1: Sends 1, 3, 5 and closes.
	go func() {
		defer close(ch1) // Close the channel when done
		fmt.Println("[Producer 1]: Sending 1")
		ch1 <- 1
		time.Sleep(100 * time.Millisecond)

		fmt.Println("[Producer 1]: Sending 3")
		ch1 <- 3
		time.Sleep(100 * time.Millisecond)

		fmt.Println("[Producer 1]: Sending 5")
		ch1 <- 5
		fmt.Println("[Producer 1]: Done, closing ch1.")
	}()

	// Producer 2: Sends 2, 4 and closes (at a different speed).
	go func() {
		defer close(ch2)                  // Close the channel when done
		time.Sleep(50 * time.Millisecond) // Start slightly later
		fmt.Println("[Producer 2]: Sending 2")
		ch2 <- 2

		time.Sleep(1000 * time.Millisecond)
		fmt.Println("[Producer 2]: Sending 4")
		ch2 <- 4
		fmt.Println("[Producer 2]: Done, closing ch2.")
	}()

	// --- Consumer (Main Goroutine) ---
	// We range over the merged channel. The loop will
	// automatically stop when 'merged' is closed.
	fmt.Println("[Main]: Waiting for merged data...")
	for v := range merged {
		fmt.Printf("[Main]: Received %d\n", v)
	}

	fmt.Println("[Main]: Merged channel closed. Program finished.")
}
