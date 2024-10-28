package functions

func BubbleSort(intSlice []int) {

	length := len(intSlice)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				Swap(intSlice, j)
			}
		}
	}

}

func Swap(slice []int, from int) {
	slice[from], slice[from+1] = slice[from+1], slice[from]
}

/*
func main() {
	fmt.Println("Bubble Sort")
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	fmt.Print("Enter up to 10 integers separated by spaces: ")
	scanner.Scan()
	input = scanner.Text()

	strNumbers := strings.Fields(input)
	if len(strNumbers) > 10 {
		fmt.Println("Error: Please enter no more than 10 integers.")
		return
	}

	intSlice := make([]int, len(strNumbers))
	for i, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error: Invalid integer input.")
			return
		}
		intSlice[i] = num
	}

	BubbleSort(intSlice)

	fmt.Println("Sorted numbers", intSlice)
}
*/
