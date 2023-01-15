package common

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file.
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}

func GenerateHash(s string) []byte {
	ss := os.Getenv("SECRETSALT")
	h := sha1.New()
	h.Write([]byte(s+ss))
	hash := h.Sum(nil)
	return hash
}