package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	animalsToTypes := make(map[string]string, 0)

	availableMethods := []string{"eat", "move", "speak"}
	types := []string{"cow", "bird", "snake"}
	fmt.Println(types)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		promptUser()
		scanner.Scan()

		input := scanner.Text()

		words := strings.Split(input, " ")
		numberOfWords := len(words)

		if numberOfWords == 3 && words[0] == "newanimal" {
			animalName := words[1]
			animalType := words[2]
			if contains_(types, animalName) {
				fmt.Println("name is reserved for type")
				continue
			}
			if !contains_(types, animalType) {
				fmt.Println("no such animal type")
				continue
			}
			animalsToTypes[animalName] = animalType
			fmt.Println("Created it!")
		} else if numberOfWords == 3 && words[0] == "query" {
			animalName := words[1]
			activityName := words[2]

			value, isMapContainsKey := animalsToTypes[animalName]

			if !isMapContainsKey {
				fmt.Println("animal with such wasn't found")
				continue
			}

			if !contains_(availableMethods, activityName) {
				fmt.Println("such method unavailable")
				continue
			}

			animal, containsInMap := mapOfAnimals[value]

			if containsInMap {

				if activityName == "eat" {
					animal.Eat()
				} else if activityName == "move" {
					animal.Move()
				} else if activityName == "speak" {
					animal.Speak()
				} else {
					continue
				}

			}

		} else {
			return
		}

	}

}
