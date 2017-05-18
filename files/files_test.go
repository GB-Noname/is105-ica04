package files

import (
"testing"
)

func benchmarkFiles (filename string, b*testing.B) {
	for file2 := 0; file2< b.N; file2++ {
		FileToByteslice(filename)
	}
}

func BenchmarkFileToByteslice(b *testing.B)  { benchmarkFiles("text1.txt",b) }




func benchmarkIoUtil(filename string, b*testing.B) {
	for dat := 0; dat<b.N; dat++{
		IoUtil(filename)
	}
}

func BenchmarkIoUtil(b *testing.B) {benchmarkIoUtil("text1.txt",b) }

func benchmarkFireader(b*testing.B) {
	for fi:=0;fi<b.N;fi++{
		Fireader()
	}

}

func   BenchmarkFireader(b *testing.B) {benchmarkFireader(b) }

