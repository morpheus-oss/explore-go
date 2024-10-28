package solved

import (
	"strings"
)

const startsWith = "i"
const endsWith = "n"
const contains = "a"

const FOUND = "Found!"
const NOT_FOUND = "Not Found!"

func Findian(str string) bool {
	str = strings.ToLower(str)

	if !strings.HasPrefix(str, startsWith) {
		return false
	}
	if !strings.HasSuffix(str, endsWith) {
		return false
	}
	if !strings.Contains(str, contains) {
		return false
	}

	return true
}

/*
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var inputString string
	fmt.Printf("Enter any string:")
	scanner.Scan()
	inputString = scanner.Text()

	if Findian(inputString) {
		fmt.Println(FOUND)
	} else {
		fmt.Println(NOT_FOUND)
	}

}
*/
