package memory

import "fmt"

func PassByValueDemo() {
	fmt.Println("-------Pass by value demo--------")
	i := 10
	fmt.Printf("Value: %d, Address: %d\n", i, &i)
	increment(i)
	fmt.Printf("Value: %d, Address: %d\n", i, &i)
	fmt.Println("---------------------------------")
}

func increment(value int) {
	value++
	fmt.Printf("Value: %d, Address: %d\n", value, &value)
}
