package main

import (
	"fmt"
	//"path/filepath"
	"log"
	"os"
	"reflect"

	"./files"
	"./lineshift"
	//"net/http"
)

func main() {

	//fmt.Printf("%s", files.FileToByteslice(text1))

	urlText1 := "https://gb-noname.github.io/text1.txt"
	urlText2 := "https://gb-noname.github.io/text2.txt"
	text1 := files.FileToByteslice(urlText1)
	text2 := files.FileToByteslice(urlText2)

	//fmt.Printf("%x\n", a)
	//fmt.Printf("%x\n", b)
	// DeepEqual tester om byter er identiske
	if reflect.DeepEqual(text1, text2) == false {
		fmt.Println("Filene er ikke identiske")
	} else {
		fmt.Println("Filene er identiske")
	}
	//fmt.Print(a)

	fmt.Println(lineshift.Tester(text1, "text1.txt"))
	fmt.Println(lineshift.Tester(text2, "text2.txt"))

<<<<<<< HEAD
	 delErr := os.Remove("temp.txt")
	 if delErr != nil {
	 log.Fatal(delErr)
	 }
=======
	delErr := os.Remove("temp.txt")
	if delErr != nil {
		log.Fatal(delErr)
	}
>>>>>>> 5c655a08334cc096ea60a7ea6788c0d85e06f73b
	//fmt.Println(reflect.DeepEqual(text1, text2))

	//fmt.Printf("% x\n", files.FileToByteslice(text1))
	//fmt.Printf("% x", files.FileToByteslice(text2))

	//fmt.Println(lineshift.Tester(a))

}
