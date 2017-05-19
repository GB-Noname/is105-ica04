package files

import (
	"io"
	"os"
)

func Fireader() {
	fi, err := os.Open("./text1.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()


	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()


	buf := make([]byte, 1024)
	for {

		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}