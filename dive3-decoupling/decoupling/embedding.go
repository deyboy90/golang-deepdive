package decoupling

import "fmt"

type person struct {
	name  string
	email string
}

func (p person) notify() {
	fmt.Printf("Notifying person, name: %s, email: %s\n", p.name, p.email)
}

type teamleader struct {
	person // embedding, all fields and methods of driver will be promoted to teamleader
	level  string
}

func (tl teamleader) notify() {
	fmt.Printf("Notifying %s, name: %s, email: %s\n", tl.level, tl.name, tl.email)
}

func EmbeddingDemo() {
	fmt.Println("-------TypeAssertionDemo--------")
	d := person{
		name:  "Ricardo",
		email: "ricardo@gmail.com",
	}

	d.notify()

	tl := teamleader{
		person{
			name:  "Christian",
			email: "christian@gmail.com",
		},
		"Captain",
	}

	tl.person.notify() // this explicitly accesses the inner type notify method
	tl.notify()        // this accesses the outer type notify method

	// Interfaces are not necessary to utilize embedding, but if we
	// want to utilize polymorphic functions then using interfaces is valuable
	sendLetter(d)
	sendLetter(tl)

	fmt.Println("---------------------------------")
}

// interfaces should always be verb vs nouns, encapsulate behavior vs state
type notifier interface {
	notify()
}

func sendLetter(n notifier) {
	fmt.Println("Sending letter ...")
	fmt.Println()
	n.notify()
	fmt.Println()
	fmt.Println("Thank You, \nTeam Owner")
}
