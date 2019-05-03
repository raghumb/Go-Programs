package main

import (
	"fmt"
	"strings"
)
 /*Validate Strong password*/
func ValidateStrongPassword(input string) (bool){
	var is_upper bool = ValidateUpperCase(input)
	var is_special bool = ValidateSpecialCharacter(input)

	fmt.Println("is_upper ",is_upper)
	fmt.Println("is_special ",is_special)

	return is_upper && is_special
 } 
 /*Validate presence of Upper characters*/
 func ValidateUpperCase(password string) (bool){
	var i int
	var leng = len(password)
	var valid bool = false
	for i=0;i < leng; i++{
		/*Validate Upper Case character*/
		var ascii = int(password[i])
		if (ascii >= int('A') && ascii <= int('Z')){
			valid = true
			break
		}

	}
	return valid
 }

 /*Validate Special characters*/
 func ValidateSpecialCharacter(input string) (bool){
	var valid bool = false
	var special_chars string = "!@#$%^&*()"
	valid = strings.ContainsAny(input, special_chars)
	return valid
 }