package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(arr []byte) (int, error) {
	count, err := rot.r.Read(arr)
	if err != nil {
		return count, err
	}
	for i, val := range arr {
		if val >= 'a' && val <= 'z' {
			arr[i] = ((val-'a')+13)%26 + 'a'
		}
		if val >= 'A' && val <= 'Z' {
			arr[i] = ((val-'A')+13)%26 + 'A'
		}
	}
	return count, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
