package helloworld_test

import (
	"fmt"
	"testing"

	helloworld "github.com/lirui311012/HelloWorld"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := helloworld.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}

}
