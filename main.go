package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	bitsFlag := flag.Int("bits", 256, "the number of bits in the secret")
	flag.Parse()
	c := *bitsFlag / 8
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	str := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("gjs:%s\n", str)
}
