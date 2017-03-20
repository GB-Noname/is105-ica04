package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("README.md") // Fil
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(b) // print i 'bytes'

	str := string(b) // convert til 'string'

	fmt.Println(str) // print som 'string'
}
