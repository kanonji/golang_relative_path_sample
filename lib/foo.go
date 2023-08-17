package lib

import (
	_ "embed"
	"log"
	"os"
)

//go:embed foo.txt
var embedded []byte

func Main() {
	log.Println("os.Getwd(): ", MustToStr(os.Getwd()))
	log.Println("os.Executable(): ", MustToStr(os.Executable()))
	log.Println("os.ReadFile(\"./foo.txt\"): ", MustToStr(os.ReadFile("./foo.txt")))
	log.Println("os.ReadFile(\"./lib/foo.txt\"): ", MustToStr(os.ReadFile("./lib/foo.txt")))
	log.Println("//go:embed foo.txt: ", string(embedded))
}

func MustToStr[T string | []byte](val T, err error) string {
	if err != nil {
		return err.Error()
	}
	return string(val)
}
