package functions

import "fmt"

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {

	return func(time float64) float64 {
		return (0.5 * acceleration * (time * time)) + (initialVelocity * time) + initialDisplacement
	}

}

func main() {
	var acceleration, initialVelocity, initialDisplacement float64
	fmt.Print("Enter acceleration (a): ")
	fmt.Scan(&acceleration)
	fmt.Print("Enter initial velocity (vo): ")
	fmt.Scan(&initialVelocity)
	fmt.Print("Enter initial displacement (so): ")
	fmt.Scan(&initialDisplacement)

	displacementfn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	var time float64
	fmt.Print("Enter time (t): ")
	fmt.Scan(&time)

	fmt.Printf("Displacement after %.2f seconds: %.2f\n", time, displacementfn(time))

}
