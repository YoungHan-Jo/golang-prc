package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

// func (f *File) Close() {
// }

func ReadFile(reader Reader) {
	if c, ok := reader.(Closer); ok {
		c.Close()
	}
}

func main() {
	file := &File{}
	ReadFile(file)
}
