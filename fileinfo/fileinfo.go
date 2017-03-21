package main

import (
<<<<<<< HEAD
	"fmt"
	"os"
=======
"fmt"
"os"
<<<<<<< HEAD
=======

>>>>>>> e69105e61acffa41b7993e917702ffa591f1230f
>>>>>>> f237e2e1961fcff448b2c95126da648346e5ef6c
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

<<<<<<< HEAD
	fmt.Println("Information about a file", arg)
=======
	fmt.Println("Information about a file")
>>>>>>> f237e2e1961fcff448b2c95126da648346e5ef6c
	fmt.Println("File size in bytes: ", bytes, "File size in kilobytes: ", kibibytes, "File size in megabytes: ", mibibytes, "File size in gigabytes: ", gibibytes)
	fmt.Println("Permissions:", file.Mode())    //..
<<<<<<< HEAD
	mode := file.Mode()
	fmt.Println("Is a directory: ", mode.IsDir())
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
=======
	fmt.Println("Is Directory: ", file.IsDir()) //..
<<<<<<< HEAD
	fmt.Println("Is regular file", file.IsRegular())

}
=======
	fmt.Printf("System Interface type: %T\n", file.Sys())
	//fmt.Printf("System info: %+v\n\n", file.Sys())
>>>>>>> f237e2e1961fcff448b2c95126da648346e5ef6c

}
>>>>>>> e69105e61acffa41b7993e917702ffa591f1230f
