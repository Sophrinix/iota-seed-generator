package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	SeedCharset = "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SeedLength  = 81
)

func GenerateRandomInt(max int64) (int64, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return bigInt.Int64(), nil
}

func GenerateRandomSeed() (string, error) {
	seed := make([]byte, SeedLength)
	for i := 0; i < SeedLength; i++ {
		n, err := GenerateRandomInt(int64(len(SeedCharset)))
		if err != nil {
			return string([]byte{}), err
		}
		seed[i] = SeedCharset[n]
	}
	return string(seed), nil
}

func main() {
	s, err := GenerateRandomSeed()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
