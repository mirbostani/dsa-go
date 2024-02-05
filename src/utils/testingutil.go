package testingutil

import "testing"

func AssertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func AssertIsNil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func AssertIsNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("Expected not nil, but got nil")
	}
}

func AssertTrue(t *testing.T, actual interface{}) {
	if actual != true {
		t.Errorf("Expected true, but got %v", actual)
	}
}
