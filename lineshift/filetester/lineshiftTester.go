package lineshift

import (
	"bytes"
	"fmt"
)

var lfIntT int = 0

func Tester(text []byte, filename string) {

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
			lfIntT++
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
	if crInt < lfInt && crInt == 0{
		buffer.WriteString("Kun LineFeed. Dette er en fil laget med/for et UNIX/OSX system.")
	} else {
		buffer.WriteString("Ulik mengde CR og LF linjeslutt! Feil i fil?")
	}
	if crInt > 0 && lfInt == 0 {
		buffer.WriteString("Kun Carriage Return! Sansyneligvis gammel fil fra classic Mac OS < v9")
		lfIntT = crInt
	}
	fmt.Println("Resultat for fil: ", filename)
	fmt.Println("Antall LF:", lfInt)
	fmt.Println("Antall CR:", crInt)

	fmt.Println( buffer.String())
	AmountOfLines()

}

func AmountOfLines() {
	fmt.Println("Det er totalt ", lfIntT, "linjer i denne filen")
	lfIntT = 0
}


