package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	fileInfo, err = os.Stat("fileinfo.go") //Velg fil ".."
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())        //..
	fmt.Println("Size in bytes:", fileInfo.Size())    //Størrelse i bytes
	fmt.Println("Permissions:", fileInfo.Mode())      //..
	fmt.Println("Last modified:", fileInfo.ModTime()) //Sist endret
	fmt.Println("Is Directory: ", fileInfo.IsDir())   //..

}
