package lineshift

import (
	"bytes"
	"fmt"
)

func Tester(text1 []byte) string {
	//text1 := "/GitHub/Org/is105-ica04/files/text1.txt"
	//var a []byte = files.FileToByteslice(text1)
	// 0a UTF8 linefeed
	lf := "0a"
	lfInt := 0
	// 0d carriage return
	cr := "0d"
	crInt := 0
	// 0c Form feed
	ff := "0c"
	ffInt := 0

	var buffer bytes.Buffer
	buffer.WriteString("Linjeendinger i teksten er: ")

	for i := 0; i < len(text1); i++ {
		fmt.Printf("%X", text1[i])
		if string(text1[i]) == lf {
			lfInt++
		}
		if string(text1[i]) == cr {
			crInt++
		}
		if string(text1[i]) == ff {
			ffInt++
		}

	}
	if crInt >= lfInt {
		buffer.WriteString("Carriage Return og LineFeed")
	}
	if crInt < lfInt {
		buffer.WriteString("Kun LineFeed")
	}
	fmt.Println("Antall LF:", lfInt)
	fmt.Println("Antall CR:", crInt)

	return buffer.String()
}
