package main

import (
	"fmt"
	"os"
)

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

	fmt.Println("Information about a file", arg)
	fmt.Println("File size in bytes: ", bytes, "File size in kibibytes: ", kibibytes, "File size in mibibytes: ", mibibytes, "File size in gibibytes: ", gibibytes)
	fmt.Println("Permissions:", file.Mode())    //..
	mode := file.Mode()
	fmt.Println("Is a directory: "	, mode.IsDir())
	fmt.Println("Is a regular file", mode.IsRegular())


	if mode&os.ModeAppend != 0 {
		fmt.Println("Is Append only file")
	} else {
		fmt.Println("Is not append only file")
	}
	if mode&os.ModeDevice != 0 {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Not a device file")
	}
	if mode&os.ModeCharDevice != 0 {
		fmt.Println("Is a UNIX character device")
	} else {
		fmt.Println("Is not a UNIX character device")
	}
	if mode&os.ModeType != 0 {
		fmt.Println("Is a UNIX block device")
	} else {
		fmt.Println("Is not a UNIX block device")
	}
	if mode&os.ModeSymlink != 0 {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}

}

