package concurrency

import (
	"sync"
)

/**
Race Condition:
A race condition occurs when two or more goroutines perform operations on a shared
variable that depend on the variable’s current state, and the goroutines interleave
their operations in an unpredictable way.

Program Explanation:
Shared Resource:
 - In this program 'sharedVar' is accessed by two goroutines: one increments the value
 	and the other decrements the value.
 - As both the goroutines are working on the same variable without syncronization, their
	operations may overlap, leading to unpredictable behavior.

Goroutines:
incrementer - increments the shared variable 1000 times.
decrementer - decrements the shared variable 1000 times.

No Race Condition:
In case if the goroutines were synchronized, the final value of 'sharedVar' would always
be '0', because 1000 increments and 1000 decrements would cancel each other out.

What happens (Race Condition):
 - Due to the lack of synchronization, both goroutines may read, modify, and write to
 	sharedVar at the same time.
 - For example, while incrementer is reading the current value of sharedVar, decrementer
	may also read the same value, causing both to update the variable based on the same
	initial value.
 - This leads to inconsistent modifications, and the final value of sharedVar will often
 	not be 0.

*/

var sharedVar int // shared variable

// First goroutine increments the shared variable
func incrementer(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		sharedVar++
	}
}

// Second goroutine decrements the shared variable
func decrementer(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		sharedVar--
	}
}

/**
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	wg.Add(1)

	// Start the first goroutine
	go incrementer(&wg)

	// Start the second goroutine
	go decrementer(&wg)

	// Wait for both goroutines to finish
	wg.Wait()

	// The final value of sharedVar should ideally be 0, but it may not be due to race conditions.
	fmt.Println("Final Value of sharedVar:", sharedVar)
}
*/
