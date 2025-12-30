package main

import (
	"errors"
	"fmt"
)

// --- 1. Base Error and Marker Error Definitions ---

// Define a sentinel error for the core failure.
var ErrPermissionDenied = errors.New("permission denied")

// Define a custom error type to serve as a MARKER.
// It also wraps the underlying error.
type ResourceAccessError struct {
	User     string
	Resource string
	Err      error // Field to hold the source error
}

// Implement the error interface for ResourceAccessError.
func (e ResourceAccessError) Error() string {
	return fmt.Sprintf("user %s failed to access resource %s: %v", e.User, e.Resource, e.Err)
}

// Implement the Wrapper interface to allow use of errors.Unwrap().
func (e ResourceAccessError) Unwrap() error {
	return e.Err
}

// Implement the Is method for checking specific types (marking the error).
// This allows us to check if the error is OUR type using errors.Is.
func (e ResourceAccessError) Is(target error) bool {
	// Check if the target error is specifically the ErrPermissionDenied sentinel.
	return errors.Is(e.Err, target)
}

// --- 2. Simulated Functions ---

// The low-level function that fails.
func accessDB() error {
	return ErrPermissionDenied
}

// --- 3. Error Handling Scenarios ---

// Scenario A: Custom Struct Wrapping (Marking + Context)
func handleScenarioA(user, resource string) error {
	err := accessDB()
	if err != nil {
		// Wrap the underlying error in our custom struct, providing context
		// (User, Resource) and marking it as a ResourceAccessError.
		return ResourceAccessError{
			User:     user,
			Resource: resource,
			Err:      err, // Source error is wrapped here
		}
	}
	return nil
}

// Scenario B: fmt.Errorf with %w (Adding Context + Wrapping)
func handleScenarioB(user string) error {
	err := accessDB()
	if err != nil {
		// Use %w to wrap the error. This adds context and makes the source
		// error available via errors.Unwrap().
		return fmt.Errorf("failed during database operation for user %s: %w", user, err)
	}
	return nil
}

// Scenario C: fmt.Errorf with %v (Adding Context + Transforming)
func handleScenarioC() error {
	err := accessDB()
	if err != nil {
		// Use %v to transform the error. This adds context but destroys
		// the availability of the source error.
		return fmt.Errorf("critical failure in processing request: %v", err)
	}
	return nil
}

// --- 4. Main Execution and Checking ---

func main() {
	// --- A. Custom Struct Wrapping Example (Marking and Context) ---
	errA := handleScenarioA("Alice", "User_Records")
	fmt.Println("--- Scenario A (Custom Struct Wrapper) ---")
	fmt.Printf("Returned Error: %v\n", errA)

	// Check if the error is the underlying sentinel error (ErrPermissionDenied)
	// errors.Is uses the custom Is method on ResourceAccessError.
	if errors.Is(errA, ErrPermissionDenied) {
		fmt.Println("  ✅ errors.Is: Source error is ErrPermissionDenied.")
	}

	// Check if the error is of the specific custom type (Marking)
	var resourceErr ResourceAccessError
	if errors.As(errA, &resourceErr) {
		fmt.Printf("  ✅ errors.As: Error is a ResourceAccessError. User: %s\n", resourceErr.User)
	}

	// --- B. %w Wrapping Example (Context and Source Available) ---
	errB := handleScenarioB("Bob")
	fmt.Println("\n--- Scenario B (fmt.Errorf with %w) ---")
	fmt.Printf("Returned Error: %v\n", errB)

	// Check if the error is the underlying sentinel error.
	if errors.Is(errB, ErrPermissionDenied) {
		fmt.Println("  ✅ errors.Is: Source error is ErrPermissionDenied.")
	}

	// Explicitly demonstrate unwrapping
	unwrappedErr := errors.Unwrap(errB)
	fmt.Printf("  Unwrapped Error: %v\n", unwrappedErr)

	// --- C. %v Transformation Example (Context but Source Lost) ---
	errC := handleScenarioC()
	fmt.Println("\n--- Scenario C (fmt.Errorf with %v) ---")
	fmt.Printf("Returned Error: %v\n", errC)

	// Check if the error is the underlying sentinel error.
	if errors.Is(errC, ErrPermissionDenied) {
		fmt.Println("  ❌ errors.Is: Source error is ErrPermissionDenied (Should Fail).")
	} else {
		fmt.Println("  ✅ errors.Is: Source error is NOT available (as expected with %v).")
	}
}
