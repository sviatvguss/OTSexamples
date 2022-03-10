package timelib

import (
	"math/rand"
	"time"
)

func GetTime() string {
	return time.Now().Format(time.RFC850)
}

func GetRandomVal(max int) int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	return rand.Intn(max) + 1
}
