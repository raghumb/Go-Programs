package main

import "testing"

func TestValidateStrongPassword(t *testing.T) {
	var result = ValidateStrongPassword("hello")
	var expected_result = true
    if result == false {
       t.Errorf("Result is incorrect, Obtained: %t, Expected: %t.", result, expected_result)
    }
}

func TestValidateUpperCase(t *testing.T) {
	var result = ValidateUpperCase("hello")
	var expected_result = true
    if result == false {
       t.Errorf("Result is incorrect, Obtained: %t, Expected: %t.", result, expected_result)
    }
}

func TestValidateSpecialCharacter(t *testing.T) {
	var result = ValidateSpecialCharacter("hello")
	var expected_result = true
    if result == false {
       t.Errorf("Result is incorrect, Obtained: %t, Expected: %t.", result, expected_result)
    }
}