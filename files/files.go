package files

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FileToByteslice(filename string) []byte {
	// her opprette TempPath fordi man ikke midlertidig kan åpne filen fra serveren, bare kopiere info fra den lul
	tempPath := "temp.txt"
	file, err := os.Create(tempPath)

	if err != nil {
		fmt.Println(err)
	}
	// Open file for reading
	url := filename
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	//open a file for writing
	file2, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	finfo, err := file2.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()


	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file2, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()
	file2.Close()
	//file.Close()

	return byteSlice

}

func WebPath (url string) []byte {
	// her opprettes TempPath fordi man ikke midlertidig kan åpne filen fra serveren, bare kopiere info fra den
	tempPath := "temp.txt"
	file, err := os.Create(tempPath)

	if err != nil {
		fmt.Println(err)
	}

	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	// Dumper response body til filen med io.Copy. Kan behandle store mengder data
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	defer file.Close()
	return FileToByteslice(tempPath)
}
