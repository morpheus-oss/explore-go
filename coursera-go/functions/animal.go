package functions

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (an Animal) Eat() {
	fmt.Println(an.food)
}

func (an Animal) Move() {
	fmt.Println(an.locomotion)
}

func (an Animal) Speak() {
	fmt.Println(an.noise)
}

var cow Animal
var bird Animal
var snake Animal

// Initializes pre-defined 'Animal' objects
func __init() {
	cow = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	bird = Animal{
		food:       "worm",
		locomotion: "fly",
		noise:      "peep",
	}
	snake = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
}

func PrintAction(animal Animal, action string) {
	switch strings.ToLower(action) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid action for the animal,", action)
		fmt.Println("Available actions: Eat, Move and Speak.")
	}
}

/*
func main() {
	__init()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()
		request := scanner.Text()
		reqParts := strings.Fields(request)

		if len(reqParts) < 2 {
			fmt.Println("Invalid request. Please provide two arguments, animal and activity.")
			continue
		}

		animal := reqParts[0]
		action := reqParts[1]

		switch strings.ToLower(animal) {
		case "cow":
			PrintAction(cow, action)
		case "bird":
			PrintAction(bird, action)
		case "snake":
			PrintAction(snake, action)
		default:
			fmt.Println("Invalid animal name", animal)
			fmt.Println("Available animals: cow, bird and snake")
		}
	}
}
*/
