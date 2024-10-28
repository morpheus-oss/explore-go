package concurrency

import (
	"fmt"
	"sort"
	"sync"
)

// Function to sort a part of the array
func sortPart(arr []int, wg *sync.WaitGroup, partName string) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Printf("Goroutine sorting %s: %v\n", partName, arr)
	sort.Ints(arr)
	fmt.Printf("Sorted %s: %v\n", partName, arr)
}

// Function to merge two sorted arrays
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Merge the two sorted arrays
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

/*
func main() {
	// Read input from the user
	fmt.Println("Enter a series of integers separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert input to an integer slice
	strNumbers := strings.Split(input, " ")
	var numbers []int
	for _, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid input. Please enter integers only.")
			return
		}
		numbers = append(numbers, num)
	}

	// Divide the array into four approximately equal parts
	length := len(numbers)
	partSize := length / 4

	// Handle uneven division
	var partitions [4][]int
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if i == 3 {
			// Make sure the last partition gets any remaining elements
			end = length
		}
		partitions[i] = numbers[start:end]
	}

	// Sort each part concurrently using goroutines
	var wg sync.WaitGroup
	wg.Add(4)

	go sortPart(partitions[0], &wg, "Part 1")
	go sortPart(partitions[1], &wg, "Part 2")
	go sortPart(partitions[2], &wg, "Part 3")
	go sortPart(partitions[3], &wg, "Part 4")

	// Wait for all goroutines to finish sorting
	wg.Wait()

	// Merge the four sorted parts
	sorted1 := merge(partitions[0], partitions[1])
	sorted2 := merge(partitions[2], partitions[3])
	finalSorted := merge(sorted1, sorted2)

	// Print the final sorted array
	fmt.Println("Final sorted array:", finalSorted)
}
*/
