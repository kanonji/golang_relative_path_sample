This is a little code to check the starting point of relative paths when loading files in Golang.
I'm checking `os.ReadFile()` and `go:embed`.

## Behaviours

* `os.ReadFile()`: the current directory obtained by `os.Getwd()` is the starting point.
* `go:embed`: starting from the file where `//go:embed` is written.

## Outputs

```
$ go run cmd/main.go
main.go:15: os.Getwd():  /path/to/golang_relative_path_sample
main.go:16: os.Executable():  /tmp/go-build3815424284/b001/exe/main
main.go:17: os.ReadFile("./foo.txt"):  open ./foo.txt: no such file or directory
main.go:18: os.ReadFile("./cmd/foo.txt"):  This is at cmd/foo.txt

main.go:19: //go:embed foo.txt:  This is at cmd/foo.txt

foo.go:13: os.Getwd():  /path/to/golang_relative_path_sample
foo.go:14: os.Executable():  /tmp/go-build3815424284/b001/exe/main
foo.go:15: os.ReadFile("./foo.txt"):  open ./foo.txt: no such file or directory
foo.go:16: os.ReadFile("./lib/foo.txt"):  This is at lib/foo.txt

foo.go:17: //go:embed foo.txt:  This is at lib/foo.txt
```
