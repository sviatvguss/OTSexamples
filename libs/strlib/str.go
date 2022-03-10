package strlib

import (
	"strings"
)

func ReplaceAtoBinC(a, b, c string) string {
	replacer := strings.NewReplacer(a, b)
	fixed := replacer.Replace(c)
	return fixed
}
