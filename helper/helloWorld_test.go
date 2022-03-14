package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Syaukhul")
	assert.Equal(t, "Hello Syaukhul", result, "Result must be 'Hello Syaukhul'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Syaukhul")
	require.Equal(t, "Hello Syaukhul", result, "Result must be 'Hello Syaukhul'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Syaukhul")
	if result != "Hello Syaukhul" {
		//error
		t.Error("Result must be 'Hello Syaukhul'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloYou(t *testing.T) {
	result := HelloWorld("You")
	if result != "Hello You" {
		//error
		t.Fatal("Result must be 'Hello You'")
	}
	fmt.Println("TestHelloYou Done")
}
