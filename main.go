// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/sfreiberg/gotwilio"
	"hash"
	"math/rand"
	"strconv"
	"time"
)

var hasher hash.Hash = sha512.New()

func main() {

	accountSid := "ACda22d76d1f0ca22dae6d9d8cff1307f6"
	authToken := "ca1704c1b5df5cdf98b564574edc69c9"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+14177202086"
	to := "+14179880783"

	// Seed the rng
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Welcome to the mega SHA512 cracker!  Let's do this thing!!")
	fmt.Println("")
	fmt.Println("")

	for {
		num := strconv.FormatInt(rand.Int63n(1800000000000000000), 10)
		h := doHash(num)
		count := countLeadingZeros(h)
		if count >= 9 {
			fmt.Printf("\nCOUNT OF %d, Source: %s\n", count, num)
			fmt.Println(h)
			twilio.SendSMS(from, to, "Got a "+strconv.FormatInt(int64(count), 10)+": "+num, "", "")
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
