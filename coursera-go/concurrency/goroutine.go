package concurrency

import (
	"sync"
)

var wgroup sync.WaitGroup

var x int = 0

func g1() {
	x = 4
	wgroup.Done()
}

func g2() {
	x = x + 1
	wgroup.Done()
}

/*
func main() {

	wgroup.Add(2)
	go g1()
	go g2()
	wgroup.Wait()
	fmt.Println(x)
}
*/
