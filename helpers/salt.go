package helpers

// File provides global helpers for salt and hashing.

import (
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSALT Generates salt based on settings value.
func GenerateSALT(length int) string {
	return RandomString(length)
}

// GenerateSessionID Generates session ID based on settings value and UNIX time.
func GenerateSessionID(length int) string {
	return UXAsSeconds() + RandomString(length)
}

// HASHPassword Uses bcrypt to generate password hash.
// Suggested hash cost is 14.
func HASHPassword(password, salt string, hashCost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(password+salt),
		hashCost,
	)
}

func UXAsSeconds() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func UXAsNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func CheckPasswordHash(password, salt, hashedPassword string) bool {
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
	return errCompare == nil
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := make([]byte, length)

	for i := range result {
		result[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(result)
}
