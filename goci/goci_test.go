package main

import (
	"testing"
)

func TestAge(t *testing.T) {

	var response string

	response = isThisTheAgeOfMyMother(32)
	if response != "no, but it's true that she doesn't look a day over 32" {
		t.Error("incorrect response")
	}

	response = isThisTheAgeOfMyMother(60)
	if response != "no, this is not the age of my mother" {
		t.Error("incorrect response")
	}

	response = isThisTheAgeOfMyMother(57)
	if response != "yes, this is the age of my mother" {
		t.Error("incorrect response")
	}
}
