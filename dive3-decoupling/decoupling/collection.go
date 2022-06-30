package decoupling

import "fmt"

type printer interface {
	print()
}

type canon struct {
	name string
}

func (c canon) print() {
	fmt.Printf("Printer Name: %s\n", c.name)
}

type epson struct {
	name string
}

func (e *epson) print() {
	fmt.Printf("Printer Name: %s\n", e.name)
}

func CollectionOfInterfacesDemo() {
	fmt.Println("-------PolymorphismDemo--------")
	c := canon{"PIXMA TR4520"}
	e := epson{"WorkForce Pro WF-3720"}

	// A slice of interface implies it's a collection of behaviors
	// choice of data semantics determine the behavior of the program
	printers := []printer{
		c,
		&e,
	}

	// When storing the data using value semantics, the change to the original value is not seen.
	// This is because a copy is stored inside the interface. When pointer semantics are used,
	// any changes to the original value are seen.
	c.name = "PROGRAF PRO-1000"
	e.name = "Home XP-4100"

	for _, p := range printers {
		p.print()
	}
	fmt.Println("---------------------------------")
}
