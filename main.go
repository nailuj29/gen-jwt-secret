package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
)

const version = "1.0.3"

func main() {
	bitsFlag := flag.Int("bits", 256, "the number of bits in the secret")
	plainFlag := flag.Bool("plain", false, "Remove the prefix from the secret")
	flag.Parse()
	c := *bitsFlag / 8
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	str := base64.StdEncoding.EncodeToString(b)
	if *plainFlag {
		fmt.Printf("%s\n", str)
	} else {
		fmt.Printf("gen-jwt-sec;%s;%d:%s\n", version, *bitsFlag, str)
	}
}
