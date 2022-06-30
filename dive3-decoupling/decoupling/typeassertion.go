package decoupling

import "fmt"

func TypeAssertionDemo() {
	fmt.Println("-------TypeAssertionDemo--------")
	myPrintln("Hello!")
	myPrintln(23)
	myPrintln(53.55423422)
	myPrintln(false)
	fmt.Println("---------------------------------")
}

// empty interface means this function can accept all kinds of data,
// it doesn't have to be a specific type nor does it have to exhibit
// specific behavior
func myPrintln(a interface{}) {
	// switch has a special functionality to allow for easier type assertion
	switch v := a.(type) {
	// %T directive prints the type of the value
	case string:
		fmt.Printf("Type: %T, Value: %s\n", v, v)
	case int:
		fmt.Printf("Type: %T, Value: %d\n", v, v)
	case float32:
		fmt.Printf("Type: %T, Value: %f\n", v, v)
	case float64:
		fmt.Printf("Type: %T, Value: %f\n", v, v)
	default:
		fmt.Printf("Type: %T, Value: %v\n", v, v)
	}
}
