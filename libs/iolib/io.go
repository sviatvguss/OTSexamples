package iolib

import (
	"bufio"
	"fmt"
	"io"
)

func Uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt
		_, err := fmt.Fprintln(output, txt)
		if err != nil {
			return err
		}
	}
	return nil
}
