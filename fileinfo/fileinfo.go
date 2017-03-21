package main

import (
	"fmt"
	"os"
)

var file os.FileInfo

func main() {
	arg := os.Args[1]

	file, _ := os.Stat(arg)

	var bytes int64
	bytes = file.Size()

	var kilobytes int64
	kilobytes = (bytes / 1024)

	var megabytes float64
	megabytes = (float64)(kilobytes / 1024) // cast to type float64

	var gigabytes float64
	gigabytes = (megabytes / 1024)

	fmt.Println("Information about a file")
	fmt.Println("File size in bytes: ", bytes, "File size in kilobytes: ", kilobytes, "File size in megabytes: ", megabytes, "File size in gigabytes: ", gigabytes)
	fmt.Println("File name:", file.Name())      //..
	fmt.Println("Permissions:", file.Mode())    //..
	fmt.Println("Is Directory: ", file.IsDir()) //..
	fmt.Println("Is regular file", file.IsRegular())

}
