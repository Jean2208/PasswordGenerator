package main

import (
	"encoding/hex"
	"math/rand"
	"time"

	"golang.org/x/crypto/argon2"
)

const ( //declaring variabales inside const makes sure these variables are not altered during program execution
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" //letters
	specialBytes   = "!@#$%^&*()-_=+[]{}\\|;':\",.<>?"                      //special characters
	numberBytes    = "0123456789"                                           //numbers
	allBytes       = letterBytes + specialBytes + numberBytes               //all bytes together
	passwordLength = 12                                                     //password length
	saltLength     = 16
	memory         = 32 * 1024
	threads        = 4
	times          = 3
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generatePassword() string {
	result := make([]byte, passwordLength)
	for i := range result {
		result[i] = allBytes[rand.Intn(len(allBytes))]
	}
	return string(result)
}

func hashingSalting(password string) []byte {
	password_to_hash := []byte(password)
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)

	if err != nil {
		panic(err)
	}

	salted_hash := argon2.IDKey(password_to_hash, salt, times, memory, threads, 32)

	return salted_hash

}

func main() {
	password := generatePassword()
	salted_hash := hex.EncodeToString(hashingSalting(password))
	println(password)
	println(salted_hash)

}
