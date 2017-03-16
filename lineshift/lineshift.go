package lineshift

import (
	"bytes"
)

func lineshiftFinder(text1 byte, text2 byte) string {
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

		if text1[i] == lf {
			lfInt++
		}
		if text1[cr] == cr {
			crInt++
		}
		if text1[ff] == ff {
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
