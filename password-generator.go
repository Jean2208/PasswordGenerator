package main

import (
	"math/rand"
	"time"
)

const ( //declaring variabales inside const makes sure these variables are not altered during program execution
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" //letters
	specialBytes   = "!@#$%^&*()-_=+[]{}\\|;':\",.<>?"                      //special characters
	numberBytes    = "0123456789"                                           //numbers
	allBytes       = letterBytes + specialBytes + numberBytes               //all bytes together
	passwordLength = 12                                                     //password length
)

func init() { //The function init runs before the main function
	rand.Seed(time.Now().UnixNano()) //The purpose of seeding a random number generator is to ensure that passwords generated are different each time the program is run
}

func generatePassword() string {
	result := make([]byte, passwordLength) //makes a slice
	for i := range result {                //for i g in range 12 in this case
		result[i] = allBytes[rand.Intn(len(allBytes))] //Assigning the values of the slice random values from the variable allBytes
	}
	return string(result) //returns the password to the main function
}

func main() {
	password := generatePassword() //calls the function to generate a password
	println(password)              //prints the password
}
