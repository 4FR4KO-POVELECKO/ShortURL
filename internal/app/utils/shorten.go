package shorten

import (
	"math/rand"
	"time"
)

const (
	upLetter   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowLetter  = "abcdefghijklmnopqrstuvwxyz"
	number     = "0123456789"
	underscore = "_"
)

func Shorten() string {
	b := make([]byte, 10)

	currentTime := time.Now().UnixNano()
	rand.Seed(currentTime)

	markPos := rand.Intn(10)

	up := 0
	low := 0
	num := 0

	for i := range b {
		if i == markPos {
			b[i] = underscore[0]
			continue
		}

		r := rand.Intn(3)

		if up == 3 {
			r = 1
		} else if low == 3 {
			r = 2
		} else if num == 3 {
			r = 0
		}

		if low == 3 && num == 3 {
			r = 0
		} else if up == 3 && num == 3 {
			r = 1
		} else if up == 3 && low == 3 {
			r = 2
		}

		if r == 0 {
			b[i] = upLetter[rand.Intn(len(upLetter))]
			up += 1
		} else if r == 1 {
			b[i] = lowLetter[rand.Intn(len(lowLetter))]
			low += 1
		} else if r == 2 {
			b[i] = number[rand.Intn(len(number))]
			num += 1
		}
	}

	return string(b)
}
