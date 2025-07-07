package a

import "testing"

func TestExample(t *testing.T) {
	// Act
	t.Log("doing something")
	// Arrange
	t.Log("setup") // want "// Arrange must appear before // Act"
	// Assert
	t.Log("checking result")
}

func TestMissingAct(t *testing.T) {
	// Arrange
	t.Log("setup")
	// Assert
	t.Log("checking result") // want "missing '// Act' section in test"
}

func TestMissingAssert(t *testing.T) {
	// Arrange
	t.Log("setup")
	// Act
	t.Log("doing something") // want "missing '// Assert' section in test"
}

func TestValidOrder(t *testing.T) {
	// Arrange
	t.Log("setup")
	// Act
	t.Log("doing something")
	// Assert
	t.Log("checking result")
}
