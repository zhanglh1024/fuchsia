package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expectedStr := "hello test demo!"
	result := hello()
	if result != expectedStr {
		t.Fatalf("Expected %s,got %S", expectedStr, result)
	}
}
