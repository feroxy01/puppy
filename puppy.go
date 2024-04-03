package puppy

import (
	"github.com/feroxy01/dog"
	"github.com/feroxy01/test"
)

func Bark() string {
	return "Woof!😀"
}
func Barks() string {
	return "Woof! Woof! Woof!"
}
func bar() string {
	return "Hallo"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}
func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
func tt() string {
	return test.Test()
}
