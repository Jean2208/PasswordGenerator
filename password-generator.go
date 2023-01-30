package main

import (
	"encoding/hex" //importing the hexadecimal encoding package for encoding byte slices to string representation
	"math/rand"    //importing the math/rand package for generating random numbers
	"time"         //importing the time package to get the current time to seed the random number generator

	"golang.org/x/crypto/argon2" //importing the argon2 password hashing library
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" //lowercase and uppercase letters
	specialBytes   = "!@#$%^&*()-_=+[]{}\\|;':\",.<>?"                      //special characters
	numberBytes    = "0123456789"                                           //all numbers
	allBytes       = letterBytes + specialBytes + numberBytes               //a concatenation of all characters
	passwordLength = 12                                                     //the length of the generated password
	saltLength     = 16                                                     //the length of the salt used in the hashing
	memory         = 32 * 1024                                              //memory used in the argon2 hashing algorithm
	threads        = 4                                                      //number of threads used in the Argon2 hashing algorithm
	times          = 3                                                      //number of iterations used in the Argon2 hashing algorithm
)

func init() {
	rand.Seed(time.Now().UnixNano()) //seeding the random number generator using the current time
}

func generatePassword() string {
	result := make([]byte, passwordLength) //creating a byte slice of length passwordLength
	for i := range result {                //iterating over the byte slice
		result[i] = allBytes[rand.Intn(len(allBytes))] //setting the value of each byte to a random character from allBytes
	}
	return string(result) //returning the password as a string
}

func hashingSalting(password string) string {
	password_to_hash := []byte(password) //converting the password string to a byte slice
	salt := make([]byte, saltLength)     //creating a byte slice of length saltLength to store the salt
	_, err := rand.Read(salt)            //generating random salt

	if err != nil {
		panic(err) //if an error occurs during salt generation, panic
	}

	hash := hex.EncodeToString(argon2.IDKey(password_to_hash, salt, times, memory, threads, 32)) //creating the hash using argon2

	return hash //returning the hash

}

func main() {
	password := generatePassword()             //generating the password
	hash := hashingSalting(password)           //obtaining the hash from the salted password and encoding the result to hexadecimal string representation
	println("Your secure password:", password) //printing the password
	println("Your SHA-256 hash:", hash)        //printing the hash

}
