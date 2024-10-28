package functions

import (
	"errors"
	"fmt"
	"strings"
)

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("Grass")
}

func (cow Cow) Move() {
	fmt.Println("Walk")
}

func (cow Cow) Speak() {
	fmt.Println("Moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("Fly")
}

func (b Bird) Speak() {
	fmt.Println("Peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("Mice")
}

func (s Snake) Move() {
	fmt.Println("Slither")
}

func (s Snake) Speak() {
	fmt.Println("Hsss")
}

type AnimalIf interface {
	Eat()
	Move()
	Speak()
}

var animalsMap = make(map[string]AnimalIf, 0)

func printAction(animal AnimalIf, action string) {
	switch strings.ToLower(action) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid action. Use eat, move and speak.")
	}
}

func getAnimal(animalType string) (AnimalIf, error) {
	var animal AnimalIf
	switch strings.ToLower(animalType) {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		return nil, errors.New("Invalid Animal type. Use cow, bird or snake.")
	}
	return animal, nil
}

/*
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()
		request := scanner.Text()
		reqParts := strings.Fields(request)

		if len(reqParts) < 2 {
			fmt.Println("Invalid command. Try newanimal or query.")
			continue
		}

		switch strings.ToLower(reqParts[0]) {
		case "newanimal":
			if len(reqParts) != 3 {
				fmt.Println("Invalid newanimal command. Expected: newanimal <name> <type>")
				continue
			}

			name := reqParts[1]
			animalType := reqParts[2]

			animal, err := getAnimal(animalType)
			if err != nil {
				fmt.Println(err)
				continue
			}
			animalsMap[name] = animal
			fmt.Println("Created it.")

		case "query":
			if len(reqParts) != 3 {
				fmt.Println("Invalid query command. Expected: query <name> <action>")
				continue
			}
			name := reqParts[1]
			animal, exists := animalsMap[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}
			action := reqParts[2]
			printAction(animal, action)

		default:
			fmt.Println("Invalid command. Try newanimal or query.")
		}
	}
}
*/
