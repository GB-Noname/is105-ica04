package main

import ("fmt"

)

const byteSlice = "\x74\x65\x73\x74\x65\x72\x20\x6c\x69\x74\x74\x20\x68\x65\x72"
//var sliceMap map[string]int

func main() {

	var sliceMap map[string]int
	sliceMap = make(map[string]int)

	for i := 0; i < len(byteSlice); i++ {
		//var charString string = "char" + string(i)
		counter := 1
		charString := string(byteSlice[i])
		get := sliceMap[charString]
		//counter ++
		if get == 0 {
			sliceMap[charString] = counter
		} else {
			number := sliceMap[charString]
			number2 := int(number)
			fmt.Println(charString)
			//fmt.Println(number2)
			//number2++
			//fmt.Println(number2)
			sliceMap[charString] = number2
		}
		//counter = 0
		fmt.Println(sliceMap)
	}
	fmt.Printf("%S", byteSlice)


}
