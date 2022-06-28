package memory

import "fmt"

func ConstantsDemo() {
	fmt.Println("-------ConstantsDemo--------")
	// When a constant is untyped, it’s considered to be of a kind. Constants of a kind can be implicitly converted by the compiler.

	// Untyped numeric constants have a precision of 256 bits as stated by the specification. They are based on a kind.
	const a = 123  // kind integer
	const f = 0.33 // kind float

	var div = a / f
	fmt.Println(div)

	const third = 1 / 3.0 // KindFloat = KindFloat(1) / KindFloat(3.0)
	fmt.Println(third)

	// nolint:staticcheck
	const zero = 1 / 3 // KindInt = KindInt(1) / KindInt(3)
	fmt.Println(zero)

	// a constant of a type promotes over a constant of a kind
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
	fmt.Println(two)
	fmt.Println("---------------------------------")
}
