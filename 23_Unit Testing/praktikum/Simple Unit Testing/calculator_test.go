package calculators

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(2, 3)

	if result == 5 {
		t.Log("Test Success")
	} else {
		t.Fatal("Test Failed!")
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(5, 2)

	if result == 3 {
		t.Log("Test Success")
	} else {
		t.Fatal("Test Failed!")
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(8, 4)

	if result == 32 {
		t.Log("Test Success")
	} else {
		t.Fatal("Test Failed!")
	}
}

func TestDivision(t *testing.T) {
	result := Division(8, 2)

	if result == 4 {
		t.Log("Test Success")
	} else {
		t.Fatal("Test Failed!")
	}
}
