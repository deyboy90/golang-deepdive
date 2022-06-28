package memory

import "fmt"

func PassByValueDemo() {
	fmt.Println("-------PassByValueDemo--------")
	i := 10
	fmt.Printf("Value: %d, Address: %d\n", i, &i)
	increment(i)
	fmt.Printf("Value: %d, Address: %d\n", i, &i)
	fmt.Println("---------------------------------")
}

func PassByPointerReferenceDemo() {
	fmt.Println("-------PassByPointerReferenceDemo--------")
	i := 10
	fmt.Printf("Value: %d, Address: %d\n", i, &i)
	incrementptr(&i)
	fmt.Printf("Value: %d, Address: %d\n", i, &i)
	fmt.Println("---------------------------------")
}

func increment(value int) {
	value++
	fmt.Printf("Value: %d, Address: %d\n", value, &value)
}

// we are still passing value not reference
// only diff is we are passing a "value of type pointer to an int"
func incrementptr(value *int) {
	*value++
	fmt.Printf("Value: %d, Address: %d\n", *value, &value)
}

//---------------------------------------------------------------

type user struct {
	name  string
	email string
}

func ReturnTypesDemo() {
	fmt.Println("-------ReturnTypesDemo--------")
	u1 := createUserV1()
	u2 := createUserV2()
	fmt.Printf("U1: %v, U2: %v \n", u1, *u2)
	fmt.Println("---------------------------------")
}

func createUserV1() user {
	u := user{
		name:  "Bob",
		email: "bob@gmail.com",
	}
	return u
}

/**
When doing static analysis the go compiler determines that the
user struct needs to be moved to the heap because it's reference
is being shared upwards in the call stack.

go build -gcflags -m=2
./pointers.go:59:2: u escapes to heap:
./pointers.go:59:2:   flow: ~r0 = &u:
./pointers.go:59:2:     from &u (address-of) at ./pointers.go:63:9
./pointers.go:59:2:     from return &u (return) at ./pointers.go:63:2
./pointers.go:59:2: moved to heap: u

**/
func createUserV2() *user {
	u := user{
		name:  "Tom",
		email: "tom@gmail.com",
	}
	return &u
}
