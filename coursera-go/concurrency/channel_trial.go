package concurrency

func mul(v1, v2 int, ch chan int) {
	ch <- v1 * v2
}

/*
func main() {
	chnl := make(chan int)
	go mul(1, 2, chnl)
	go mul(3, 4, chnl)

	p1 := <-chnl
	p2 := <-chnl
	fmt.Println("P1:", p1, ", p2:", p2)
}
*/
