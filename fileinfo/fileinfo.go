package main

import (
"fmt"

"os"

)

type FileMode uint32


func main() {
	arg := os.Args[1]

	file, _ := os.Stat(arg)


	var bytes int64
	bytes = file.Size()

	var kibibytes int64
	kibibytes = (bytes / 1024)

	var mibibytes float64
	mibibytes = (float64)(kibibytes / 1024) // cast to type float64

	var gibibytes float64
	gibibytes = (mibibytes / 1024)


	fmt.Println("Information about a file")
	fmt.Println("File size in bytes: ", bytes, "File size in kilobytes: ", kibibytes, "File size in megabytes: ", mibibytes, "File size in gigabytes: ", gibibytes)
	fmt.Println("File name:", file.Name())      //..
	fmt.Println("Permissions:", file.Mode())    //..
	fmt.Println("Is Directory: ", file.IsDir()) //..
	fmt.Printf("System Interface type: %T\n", file.Sys())
	//fmt.Printf("System info: %+v\n\n", file.Sys())

}