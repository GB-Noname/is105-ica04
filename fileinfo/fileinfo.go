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

	var terabytes float64
	terabytes = (gigabytes / 1024)

	var petabytes float64
	petabytes = (terabytes / 1024)

	var exabytes float64
	exabytes = (petabytes / 1024)

	var zettabytes float64
	zettabytes = (exabytes / 1024)

	fmt.Println("File name:", file.Name())        //..
	fmt.Println("Size in bytes:", file.Size())    //Størrelse i bytes
	fmt.Println("Permissions:", file.Mode())      //..
	fmt.Println("Last modified:", file.ModTime()) //Sist endret
	fmt.Println("Is Directory: ", file.IsDir())   //..


	fmt.Println("File size in bytes ", bytes)
	fmt.Println("File size in kilobytes ", kilobytes)
	fmt.Println("File size in megabytes ", megabytes)
	fmt.Println("File size in gigabytes ", gigabytes)
	fmt.Println("File size in terabytes ", terabytes)
	fmt.Println("File size in petabytes ", petabytes)
	fmt.Println("File size in exabytes ", exabytes)
	fmt.Println("File size in zettabytes ", zettabytes)


}
