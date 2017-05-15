package main

import (
	"../files"
	"./filetester"
	"os"
)

func main() {
	arg := os.Args[1]

	text := files.FileToByteslice(arg)
	lineshift.Tester(text, arg)
}
