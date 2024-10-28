package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	numEats         = 3
	maxConcurrent   = 2
)

type Philosopher struct {
	id        int
	leftChop  *Chopstick
	rightChop *Chopstick
	eatCount  int
}

type Chopstick struct {
	id   int
	lock sync.Mutex
}

type Host struct {
	mu     sync.Mutex
	active int
}

func (h *Host) allowEating() bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.active < maxConcurrent {
		h.active++
		return true
	}
	return false
}

func (h *Host) doneEating() {
	h.mu.Lock()
	h.active--
	h.mu.Unlock()
}

func (p *Philosopher) eat(host *Host) {
	for p.eatCount < numEats {
		// Simulate thinking time
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

		if host.allowEating() {
			// Pick up chopsticks in random order
			if rand.Intn(2) == 0 {
				p.leftChop.lock.Lock()
				p.rightChop.lock.Lock()
			} else {
				p.rightChop.lock.Lock()
				p.leftChop.lock.Lock()
			}

			fmt.Printf("starting to eat %d\n", p.id)
			p.eatCount++

			// Simulate eating time
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))

			fmt.Printf("finishing eating %d\n", p.id)

			p.rightChop.lock.Unlock()
			p.leftChop.lock.Unlock()
			host.doneEating()
		}
	}
}

/*
func main() {
	rand.Seed(time.Now().UnixNano())

	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := range chopsticks {
		chopsticks[i] = &Chopstick{id: i}
	}

	host := &Host{}
	var wg sync.WaitGroup

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		philosopher := &Philosopher{
			id:        i + 1,
			leftChop:  chopsticks[i],
			rightChop: chopsticks[(i+1)%numPhilosophers],
		}
		go func(p *Philosopher) {
			defer wg.Done()
			p.eat(host)
		}(philosopher)
	}

	wg.Wait()
}
*/
