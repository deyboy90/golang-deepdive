package decoupling

import "fmt"

// interfaces are valueless types
type reader interface {
	read(bytes []byte) (int, error)

	// this is not a good api design because it forces creation of data
	// on the heap, anytime we have a function that returns a slice
	// the compiler cannot store that in stack as slices are
	// backed by pointers
	//read(i int) ([]byte, error)
}

type file struct {
	name string
}

// the concrete type file now implements the reader interface using value semantics
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

type pipe struct {
	name string
}

// the concrete type pipe now implements the reader interface using value semantics
func (pipe) read(b []byte) (int, error) {
	s := `{name: "Bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

type zip struct {
	name string
}

// the concrete type zip now implements the reader interface using pointer semantics
func (*zip) read(b []byte) (int, error) {
	s := `zip contents`
	copy(b, s)
	return len(s), nil
}

// function retrieve is a polymorphic function because it’s asking for concrete data
// not based on what the data is (concrete type), but based on what the data can do (interface type).
func retreive(r reader) error {
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

func PolymorphismDemo() {
	fmt.Println("-------PolymorphismDemo--------")
	f := file{
		name: "meh.txt",
	}
	_ = retreive(f)

	p := pipe{
		name: "meh2.txt",
	}
	_ = retreive(p)

	z := zip{
		name: "meh3.txt",
	}

	// cannot use z (variable of type zip) as reader value in argument to retreive:
	// zip does not implement reader (method read has pointer receiver)compilerInvalidIfaceAssign
	// the compiler cannot determine for sure if all values will have pointers when making this call
	// hence it complaints...
	// retreive(z)

	// this works because we are explicitly specifying the pointer
	_ = retreive(&z)

	/**
	The compiler is telling me that I am not allowed to use value semantics if I have chosen to use pointer semantics.
	In other words, I am being forced to share the value with the interface since it’s not safe to make a copy of a
	value that a pointer points to. If I chose to implement the method with pointer semantics, I am stating that a
	value of this type isn’t safe to be copied.
	**/
	fmt.Println("---------------------------------")
}
