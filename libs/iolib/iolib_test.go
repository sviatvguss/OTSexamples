package iolib

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test for OK Failed - results not match\n %v %v",
			result, testOkResult)
	}
}
