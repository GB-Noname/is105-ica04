package main

import (
	"fmt"
	"path/filepath"
	"reflect"

	"./files"
	"./lineshift"
)

func main() {

	//files.FileToByteslice()
	//
	text1, _ := filepath.Abs("/Users/Ann Margrethe/Documents/GitHub/GB-noname/is105-ica04/files/text1.txt")
	fmt.Printf("%s", files.FileToByteslice(text1))
	a := files.FileToByteslice(text1)
	//
	text2, _ := filepath.Abs("/Users/Ann Margrethe/Documents/GitHub/GB-noname/is105-ica04/files/text2.txt")
	fmt.Printf("%s", files.FileToByteslice(text2))
	b := files.FileToByteslice(text2)

	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(text1, text2))

	fmt.Printf("%x\n", files.FileToByteslice(text1))
	fmt.Printf("%x", files.FileToByteslice(text2))

	lineshift.lineshiftFinder(a, b)
}
