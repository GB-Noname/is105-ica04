package files

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
	"io"
)
//var filename string = "../temp.txt"
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IoUtil (filename string) {
	//filename = file

	dat, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Print(string(dat))
	ioutil.WriteFile("temp.txt", dat, 0644)
	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.

}

func Buffio(filename string) {
	f, err := os.Open(filename)
	check(err)
	r4 := bufio.NewReader(f)

	b4, err := r4.Peek(5)
	//check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r4.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		f.Close()
	}

}


