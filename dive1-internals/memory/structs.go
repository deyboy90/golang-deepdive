package memory

import "fmt"

type example struct {
	flag    bool
	counter int32
	pi      float32
}

type anotherexample struct {
	flag    bool
	counter int32
	pi      float32
}

func StructDeclarationTypesDemo() {
	fmt.Println("-------StructDeclarationTypesDemo--------")
	// this is a named type
	e1 := example{
		flag:    true,
		counter: int32(1000),
		pi:      float32(3.14),
	}
	fmt.Println(e1)

	// this is a named type
	e2 := anotherexample{
		flag:    true,
		counter: int32(2000),
		pi:      float32(3.14),
	}
	fmt.Println(e2)

	// cannot use e2 (variable of type anotherexample) as example value in assignment compiler IncompatibleAssignment
	// e1 = e2

	// this is an anonymous type
	e3 := struct {
		flag    bool
		counter int32
		pi      float32
	}{
		flag:    true,
		counter: int32(3000),
		pi:      float32(3.14),
	}

	fmt.Println(e3)

	// this assignment is allowed because the compiler allows anonymous type value to be assigned to a named type if
	// the structure is identitcal and compatible
	e1 = e3
	fmt.Println(e1)
	fmt.Println("---------------------------------")

}
