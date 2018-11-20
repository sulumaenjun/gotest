package main

import (
	"log"
	"encoding/hex"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	input := "abcdefgh"
	//SHA1算法
	hashInBytes := sha256.Sum256([]byte(input))
	hashInStr := hex.EncodeToString(hashInBytes[:])
	log.Printf("formart:%s, %s",input,hashInStr)
	inputz := "11111"
	//SHA256算法
	hashInBytesz := sha256.Sum256([]byte(inputz))
	hashInStr1 := hex.EncodeToString(hashInBytesz[:])
	log.Printf("formart:%s, %s",inputz,hashInStr1)
}