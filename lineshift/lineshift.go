package lineshift

import (
	"bytes"
	"fmt"
)

func Tester(text1 []byte) string {
	//text1 := "/GitHub/Org/is105-ica04/files/text1.txt"
	//var a []byte = files.FileToByteslice(text1)
	// 0a UTF8 linefeed
	lf := 10
	lfInt := 0
	// 0d carriage return
	cr := 13
	crInt := 0
	// 0c Form feed
	ff := 12
	ffInt := 0


	var buffer bytes.Buffer
	buffer.WriteString("Linjeendinger i teksten er: ")

	for i := 0; i < len(text1); i++ {
		//fmt.Printf("%U", text1[i])
		//fmt.Println(text1[i])
		if int(text1[i]) == lf {
			lfInt++
		}
		if int(text1[i]) == cr {
			crInt++
		}
		if int(text1[i]) == ff {
			ffInt++
		}

	}
	if crInt == lfInt {
		buffer.WriteString("Carriage Return og LineFeed.")
		buffer.WriteString("Dette betyr at det er fil laget med Windows/DOS system.")
	}
	if crInt < lfInt {
		buffer.WriteString("Kun LineFeed. Dette er en fil laget med UNIX/OSX system.")
	}
	fmt.Println("Antall LF:", lfInt)
	fmt.Println("Antall CR:", crInt)

	return buffer.String()
}
