package files

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
	"io"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IoUtil (filename string) {


	dat, err := ioutil.ReadFile(filename)
	check(err)

	ioutil.WriteFile("temp.txt", dat, 0644)


}

func Buffio(filename string) {
	f, err := os.Open(filename)
	check(err)
	r4 := bufio.NewReader(f)

	b4, err := r4.Peek(5)

	fmt.Printf("5 bytes: %s\n", string(b4))
	buf := make([]byte, 1024)
	for {

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


