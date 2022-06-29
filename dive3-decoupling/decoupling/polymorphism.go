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
	fmt.Println("---------------------------------")
}
