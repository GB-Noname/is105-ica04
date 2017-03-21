package main

import ("fmt"
	"os"
	"log"
	)

func main () {
	arg := os.Args[1]
	file, err := os.Stat(arg)
	if err != nil {
		log.Fatal(err)
	}
	if os.ModeDevice == 0 {
		fmt.Println("Not a device file")
	}
		switch mode := file.Mode(); {
		case mode.IsDir():
			fmt.Println("directory")
		case mode.IsRegular():
			fmt.Println("regular file")
		case mode&os.ModeAppend != 0:
			fmt.Println("Append only file")
		case mode&os.ModeDevice != 0:
			fmt.Println("Is a device file")
		case mode&os.ModeCharDevice != 0:
			fmt.Println("Is a UNIX character device")
		case mode&os.ModeType != 0:
			fmt.Println("Is a UNIX block device")
		case mode&os.ModeSymlink != 0:
			fmt.Println("symbolic link")
		}

}

