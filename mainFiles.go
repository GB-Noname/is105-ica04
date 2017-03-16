package main

import (
	"fmt"
	//"path/filepath"
	"reflect"

	"./files"
	//"./lineshift"
	//"net/http"
)

func main() {

	//files.FileToByteslice()
	//
	//text1, _ := filepath.Abs("//github.com/uia-worker/is105-ica04/tree/master/files/text1.txt")

	//text1 := http.Get("https://gb-noname.github.io/text1.txt")
	//text1 := "http://gb-noname.github.io/text1.txt"
	//fmt.Printf("%s", files.FileToByteslice(text1))
	//a := files.FileToByteslice("https://gb-noname.github.io/text1.txt")

	urlText1 := "https://gb-noname.github.io/text1.txt"
	urlText2 := "https://gb-noname.github.io/text2.txt"
	a := files.FileToByteslice(urlText1)
	b := files.FileToByteslice(urlText2)


	//text2, _ := filepath.Abs("httpS//gb-noname.github.io/text1.txt")
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
	//b := files.FileToByteslice(text2)

	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Printf("%s", a)
	//fmt.Println(reflect.DeepEqual(text1, text2))

	//fmt.Printf("% x\n", files.FileToByteslice(text1))
	//fmt.Printf("% x", files.FileToByteslice(text2))

	//fmt.Println(lineshift.Tester(a))

}
