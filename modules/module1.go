package module1

import "fmt"

func Greet(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
