package timelib

import "time"

func GetTime() string {
	return time.Now().Format(time.RFC850)
}
