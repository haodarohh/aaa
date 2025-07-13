package a

import "testing"

func TestSomethingGood(t *testing.T) {
	// arrange
	// act
	// assert
}

func TestAnotherGoodOne(t *testing.T) {
	// act
	// assert
}

func TestMissingAssert(t *testing.T) { // want `missing required keywords: need at least act and assert\n`
	// arrange
	// act
}

func TestWrongOrder(t *testing.T) { //  want `invalid AAA pattern order: must follow arrange -> act -> assert\n`
	// assert
	// act
}

func TestWrongCase(t *testing.T) { // want `missing required keywords: need at least act and assert\n`
	// Arrange
	// Act
	// Assert
}

func TestNoComments(t *testing.T) { // want `missing required keywords: need at least act and assert\n`
}
