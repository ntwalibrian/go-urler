package api

import (
	"math/rand"
)

// function to shorten the url 

// function to resolve url issues

// FUNCTION TO generate short key


func RandKey(n int) string{
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}