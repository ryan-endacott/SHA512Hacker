// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"math/rand"
	"strconv"
	"time"
)

var hasher hash.Hash = sha512.New()

func main() {

	// Seed the rng
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Welcome to the mega SHA512 cracker!  Let's do this thing!!")
	fmt.Println("")
	fmt.Println("")

	for {
		num := strconv.FormatInt(rand.Int63n(1800000000000000000), 10)
		h := doHash(num)
		count := countLeadingZeros(h)
		if count > 6 {
			fmt.Printf("\nCOUNT OF %d, Source: %s\n", count, num)
			fmt.Println(h)
		}
	}
}

func doHash(str string) string {
	hasher.Reset()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func countLeadingZeros(str string) int {
	zero := byte('0')
	zeroCount := 0
	for i := 0; i < 64; i++ {
		if str[i] == zero {
			zeroCount += 1
		} else {
			return zeroCount
		}
	}
	return zeroCount
}
