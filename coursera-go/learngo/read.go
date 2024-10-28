package solved

/*
import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fName string
	lName string
}

func main() {

	var fileName string

	fmt.Print("Enter full path of the file:")
	fmt.Scan(&fileName)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error in opening file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var names []Name

	for _, line := range lines {
		nameParts := strings.Split(line, " ")

		if len(nameParts) != 2 {
			fmt.Println("Invalid name format:", nameParts)
		}

		name := Name{
			fName: strings.TrimSpace(nameParts[0]),
			lName: strings.TrimSpace(nameParts[1]),
		}

		names = append(names, name)
	}

	for _, name := range names {
		fmt.Println("First Name:", name.fName, ", Last Name:", name.lName)
	}
}
*/
