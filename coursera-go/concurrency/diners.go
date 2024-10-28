/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself,
where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself,
where <number> is the number of the philosopher.
*/

package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	func main() {
		var wg sync.WaitGroup
		channels := [5]chan int{
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
		}
		wg.Add(len(channels) + 1)

		const eatTimes = 3
		for id, channel := range channels {
			go philosopher(&wg, id, channel, eatTimes)
		}
		go host(&wg, channels[:])

		wg.Wait()
	}
*/
const REQUEST_EAT = 1
const ANSWER_NOT_ALLOWED = -1
const ANSWER_ALLOWED = -2
const NOTIFY_FINISHED_ROUND = 2

func philosopher(wg *sync.WaitGroup, philosopherId int, channel chan int, eatTimes int) {
	eatenTimes := 0
	for eatenTimes < eatTimes {
		channel <- REQUEST_EAT

		answer := <-channel
		if answer == ANSWER_ALLOWED {
			eatenTimes += 1
			fmt.Println("starting to eat [", philosopherId, "]. Round: [", eatenTimes, "]")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("finishing eating [", philosopherId, "]. Round: [", eatenTimes, "]")
			channel <- NOTIFY_FINISHED_ROUND
		} else if answer == ANSWER_NOT_ALLOWED {
			runtime.Gosched()
		}
	}
	close(channel)
	wg.Done()
}

func host(wg *sync.WaitGroup, channels []chan int) {
	numOfEaters := len(channels)
	done := 0
	philosophersEating := make([]bool, len(channels))
	philosophersFinished := make([]bool, len(channels))
	for done < numOfEaters {
		for philosopherId, channel := range channels {
			if philosophersFinished[philosopherId] {
				continue
			}
			if done == numOfEaters {
				break
			} else if philosophersFinished[philosopherId] {
				continue
			}
			select {
			case msg, ok := <-channel:
				if !ok {
					done += 1
					philosophersFinished[philosopherId] = true
					fmt.Println("Eater [", philosopherId, "] has finished eating ALL. EATING:", philosophersEating, " - FINISHED:", philosophersFinished, "DONE:", done)
				} else if msg == NOTIFY_FINISHED_ROUND {
					philosophersEating[philosopherId] = false
					fmt.Println("Eater [", philosopherId, "] has finished eating round. EATING:", philosophersEating, " - FINISHED:", philosophersFinished, "DONE:", done)
				} else if msg == REQUEST_EAT {
					left := mod_circular(philosopherId-1, numOfEaters)
					right := mod_circular(philosopherId+1, numOfEaters)
					if philosophersEating[left] || philosophersEating[right] {
						channel <- ANSWER_NOT_ALLOWED
					} else {
						philosophersEating[philosopherId] = true
						fmt.Println("Eater [", philosopherId, "] ALLOWED. EATING:", philosophersEating, " - FINISHED:", philosophersFinished, "DONE:", done)
						channel <- ANSWER_ALLOWED
					}
				}
			default:
				runtime.Gosched()
			}
		}
	}
	wg.Done()
}

// fmt.Println(mod(-0, 4))
// fmt.Println(mod(-1, 4))
// fmt.Println(mod(-2, 4))
// fmt.Println(mod(-3, 4))
// fmt.Println(mod(-4, 4))
// fmt.Println(mod(-5, 4))
// fmt.Println(mod(-6, 4))
// fmt.Println(mod(-7, 4))
// fmt.Println(mod(-8, 4))
// fmt.Println(mod(-9, 4))
// fmt.Println(mod(-10, 4))
// fmt.Println(mod(-11, 4))
// fmt.Println(mod(-12, 4))
// fmt.Println(mod(0, 4))
// fmt.Println(mod(1, 4))
// fmt.Println(mod(2, 4))
// fmt.Println(mod(3, 4))
// fmt.Println(mod(4, 4))
// fmt.Println(mod(5, 4))
// fmt.Println(mod(6, 4))
// fmt.Println(mod(7, 4))
// fmt.Println(mod(8, 4))
// fmt.Println(mod(9, 4))
// fmt.Println(mod(10, 4))
// fmt.Println(mod(11, 4))
// fmt.Println(mod(12, 4))
func mod_circular(index int, max int) int {
	if index < 0 {
		return mod_circular(max+index, max)
	} else {
		return index % max
	}
}
