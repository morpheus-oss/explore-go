package concurrency

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func __init() {
	fmt.Println("init")
}

func printStuff() {
	once.Do(__init)
	fmt.Println("hello")
	wg.Done()
}

/*
func main() {

	wg.Add(2)
	go printStuff()
	go printStuff()
	wg.Wait()

}
*/
