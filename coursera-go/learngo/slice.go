package solved

/*
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	intSlice := make([]int, 0)

	for {
		fmt.Printf("Enter a number (or 'X' to exit):")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "X" {
			break
		}

		num, error := strconv.Atoi(input)
		if error != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to exit.")
			continue
		}

		intSlice = append(intSlice, num)
		sort.Ints(intSlice)

		fmt.Println("Sorted Slice:", intSlice)
	}
}
*/
