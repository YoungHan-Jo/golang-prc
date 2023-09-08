package main

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

type Sample struct {
}

func (s *Sample) String() {
}

func CheckAndRun(stringer Stringer) {

	if r, isOk := stringer.(Reader); isOk {
		r.Read()
	}

	// r := stringer.(Reader)
	// r.Read()
}

func main() {
	sample := &Sample{}
	CheckAndRun(sample)
}
