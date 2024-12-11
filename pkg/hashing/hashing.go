package hashing

import(
    "crypto/sha256"
    "fmt"
)

func HashPassword(password string) (string) {

    sum := sha256.Sum256([]byte(password))

    hashWord := fmt.Sprintf("%x", sum)

	return hashWord
}
