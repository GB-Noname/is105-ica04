package main

import (
	"fmt"
	//"path/filepath"
	"reflect"
	"os"
	"log"
	"./files"
	"./lineshift"
	//"net/http"
)

func main() {

	//fmt.Printf("%s", files.FileToByteslice(text1))

	urlText1 := "https://gb-noname.github.io/text1.txt"
	urlText2 := "https://gb-noname.github.io/text2.txt"
	text1:= files.FileToByteslice(urlText1)
	text2 := files.FileToByteslice(urlText2)

	//fmt.Printf("%x\n", a)
	//fmt.Printf("%x\n", b)

	if reflect.DeepEqual(text1, text2) == false {
		fmt.Println("Filene er ikke identiske")
	} else {
		fmt.Println("Filene er identiske")
	}
	//fmt.Print(a)

	fmt.Println(lineshift.Tester(text1))
	fmt.Println(lineshift.Tester(text2))

	 delerr := os.Remove("temp.txt")
	 if delerr != nil {
	 log.Fatal(delerr)
	 }
	//fmt.Println(reflect.DeepEqual(text1, text2))

	//fmt.Printf("% x\n", files.FileToByteslice(text1))
	//fmt.Printf("% x", files.FileToByteslice(text2))

	//fmt.Println(lineshift.Tester(a))

}
