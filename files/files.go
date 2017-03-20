package files

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FileToByteslice(filename string) []byte {
	// her opprette TempPath fordi man ikke midlertidig kan åpne filen fra serveren, bare kopiere info fra den
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
	file2, err := os.Open(tempPath)
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	finfo, err := file2.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()
	defer response.Body.Close()

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
	file.Close()

	return byteSlice
//the laser is ready to fire
}
