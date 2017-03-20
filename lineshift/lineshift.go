package lineshift

import (
	"bytes"
)

func LineshiftFinder(text1 byte, text2 byte) string {
	//text1, _ := filepath.Abs("/Users/Ann Margrethe/Documents/GitHub/GB-noname/is105-ica04/files/text1.txt")
	//a := files.FileToByteslice(text1)

	// 0a linefeed UTF-8
	lf := "0a"
	lfInt := 0
	// 0d carriage return
	cr := "0d"
	crInt := 0
	// 0c form feed
	ff := "0c"
	ffInt := 0

	var buffer bytes.Buffer
	buffer.WriteString("Linjeshift i teksten er:")

	for i := 0; i < len(text1); i++ {

		if text1 == lf {
			lfInt++
		}
		if text1(i) == cr {
			crInt++
		}
		if text1(i) == ff {
			ffInt++
		}

		if crInt >= lfInt {
			buffer.WriteString("Carriage Return og LineFeed")
		}
		if crInt < lfInt {
			buffer.WriteString("Kun LineFeed")
		}

		return buffer.String()
	}

}
