package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	bits = 256
)

func main() {
	c := bits / 8
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	str := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("jwtSec:%s\n", str)
}
