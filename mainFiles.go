package main

import (
	"fmt"
	//"path/filepath"
	"reflect"

	"./files"
	"./lineshift"
	//"net/http"
)

func main() {

	//fmt.Printf("%s", files.FileToByteslice(text1))

	urlText1 := "https://gb-noname.github.io/text1.txt"
	urlText2 := "https://gb-noname.github.io/text2.txt"
	a := files.FileToByteslice(urlText1)
	b := files.FileToByteslice(urlText2)

	//fmt.Printf("%x\n", a)
	//fmt.Printf("%x\n", b)

	if reflect.DeepEqual(a, b) == false {
		fmt.Println("Filene er ikke identiske")
	} else {
		fmt.Println("Filene er identiske")
	}
	//fmt.Print(a)

	fmt.Println(lineshift.Tester(a))
	fmt.Println(lineshift.Tester(b))
	//fmt.Println(reflect.DeepEqual(text1, text2))

	//fmt.Printf("% x\n", files.FileToByteslice(text1))
	//fmt.Printf("% x", files.FileToByteslice(text2))

	//fmt.Println(lineshift.Tester(a))

}
