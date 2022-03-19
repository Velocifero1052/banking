package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalType struct {
	food       string
	locomotion string
	sound      string
}

func (animal AnimalType) Eat() {
	fmt.Println(animal.food)
}

func (animal AnimalType) Move() {
	fmt.Println(animal.locomotion)
}

func (animal AnimalType) Speak() {
	fmt.Println(animal.sound)
}

func promptUser() {
	fmt.Print(">")
}

func promptMethods(availableMethods []string) {
	fmt.Println("Choose what you want to know: ")
	for _, value := range availableMethods {
		fmt.Println(value)
	}
}

func contains_(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func main() {

	mapOfAnimals := map[string]AnimalType{"cow": {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"}}

	availableMethods := []string{"locomotion", "eat", "noise"}

	promptUser()

	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}

		if input == "exit" {
			return
		}

		animal, contains := mapOfAnimals[input]
		if !contains {
			fmt.Println("no such animal")
			return
		}

		promptMethods(availableMethods)

		_, err = fmt.Scan(&input)

		if err != nil {
			return
		}

		if !contains_(availableMethods, input) && input != "exit" {
			fmt.Println("no such method")
			return
		} else if input == "exit" {
			fmt.Println("Good bye!")
			return
		}

		if input == "eat" {
			animal.Eat()
		} else if input == "locomotion" {
			animal.Move()
		} else {
			animal.Speak()
		}
		promptUser()
	}

}
