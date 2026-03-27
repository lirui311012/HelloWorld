package helloworld

import "fmt"

// Hello returns hello message
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("hello %s!", name)
}
