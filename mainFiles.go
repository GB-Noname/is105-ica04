package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"./files"
	"./lineshift/filetester"

)
func main() {

	urlText1 := "https://gb-noname.github.io/text1.txt"
	urlText2 := "https://gb-noname.github.io/text2.txt"
	text1 := files.WebPath(urlText1)
	text2 := files.WebPath(urlText2)

	// DeepEqual tester om byter er identiske
	if reflect.DeepEqual(text1, text2) == false {
		fmt.Println("Filene er ikke identiske")
	} else {
		fmt.Println("Filene er identiske")
	}

	lineshift.Tester(text1, "text1.txt")

	lineshift.Tester(text2, "text2.txt")
	//TesterWithArgs()
	bigfile := files.FileToByteslice("files/pg100.txt")

	lineshift.FileStatCounter(bigfile)
	lineshift.Tester(bigfile,"files/pg100.txt" )

	files.IoUtil("files/text1.txt")

	 delErr := os.Remove("temp.txt")
	 if delErr != nil {
	 log.Fatal(delErr)
	 }
}

