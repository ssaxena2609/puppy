package puppy

import "github.com/ssaxena2609/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBark(s string) string {
	return dog.WhenGrownUp(Bark())
}
func BigBarks(s string) string {
	return dog.WhenGrownUp(Barks())
}
