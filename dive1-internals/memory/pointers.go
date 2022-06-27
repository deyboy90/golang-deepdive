package memory

import "fmt"

func PassByValueDemo() {
	i := 10
	fmt.Printf("Value: %d, Address: %d\n", i, &i)
	increment(i)
	fmt.Printf("Value: %d, Address: %d\n", i, &i)

}

func increment(value int) {
	value++
	fmt.Printf("Value: %d, Address: %d\n", value, &value)
}
