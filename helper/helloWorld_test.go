package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Andi",
			request:  "Andi",
			expected: "Hello Andi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Syaukhul", func(t *testing.T) {
		result := HelloWorld("Syaukhul")
		require.Equal(t, "Hello Syaukhul", result, "Result must be 'Hello Syaukhul'")
	})
	t.Run("A'la", func(t *testing.T) {
		result := HelloWorld("A'la")
		require.Equal(t, "Hello A'la", result, "Result must be 'Hello A'la'")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Syaukhul")
	require.Equal(t, "Hello Syaukhul", result, "Result must be 'Hello Syaukhul'")
}

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
