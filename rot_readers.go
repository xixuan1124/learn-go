package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r rot13Reader) Read(b []byte) (int, error){
	n, err := rot13r.r.Read(b)
	for i := 0; i < n; i++{
		v := &b[i]
		switch {
		case *v >= 'a' && *v <= 'z':
			*v = 'a'+(*v - 'a' + 13) % 26
		case *v >= 'A' && *v <= 'Z':
			*v = 'A' + (*v - 'A' + 13) % 26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
