package api

import (
	"encoding/base64"
	"math/rand"
	"time"
)

// validName matches a valid name string.
// var validName = regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)
// randId returns a string of random letters.

func RandId(l int) string {
	n := make([]byte, l)
	for i := range n {
		n[i] = 'a' + byte(rand.Intn(26))
	}
	return base64.URLEncoding.EncodeToString(n)
}

func init() {
	// Seed number generator with the current time.
	rand.Seed(time.Now().UnixNano())
}

func EncodeSecret(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}
