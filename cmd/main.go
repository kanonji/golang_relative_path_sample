package main

import (
	_ "embed"
	"golang_relative_path_sample/lib"
	"log"
	"os"
)

//go:embed foo.txt
var embedded []byte

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("os.Getwd(): ", lib.MustToStr(os.Getwd()))
	log.Println("os.Executable(): ", lib.MustToStr(os.Executable()))
	log.Println("os.ReadFile(\"./foo.txt\"): ", lib.MustToStr(os.ReadFile("./foo.txt")))
	log.Println("os.ReadFile(\"./cmd/foo.txt\"): ", lib.MustToStr(os.ReadFile("./cmd/foo.txt")))
	log.Println("//go:embed foo.txt: ", string(embedded))

	lib.Main()
}
