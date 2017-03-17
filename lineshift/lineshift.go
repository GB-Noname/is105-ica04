package lineshift

import (
	"bytes"
	"fmt"
)

func Tester(text []byte, filename string) string {

	//  Unicode linefeed
	lf := 10
	lfInt := 0
	//  Unicode carriage return
	cr := 13
	crInt := 0
	// Unicode Form feed
	ff := 12
	ffInt := 0


	var buffer bytes.Buffer
	buffer.WriteString("Linjeendinger i teksten er: ")

	for i := 0; i < len(text); i++ {
		//fmt.Printf("%U", text1[i])
		//fmt.Println(text1[i])
		if int(text[i]) == lf {
			lfInt++
		}
		if int(text[i]) == cr {
			crInt++
		}
		if int(text[i]) == ff {
			ffInt++
		}

	}
	if crInt == lfInt {
		buffer.WriteString("Carriage Return og LineFeed.")
		buffer.WriteString("Dette betyr at det er fil laget med/for et Windows/DOS system.")
	}
	if crInt < lfInt {
		buffer.WriteString("Kun LineFeed. Dette er en fil laget med/for et UNIX/OSX system.")
	}
	if ffInt > 0 {
		buffer.WriteString("Hvorfor er det form feed her? o.O")
	}
	fmt.Println("Resultat for fil: ", filename)
	fmt.Println("Antall LF:", lfInt)
	fmt.Println("Antall CR:", crInt)

	return buffer.String()
}
