package decoupling

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("Changed User Email To %s\n", email)
}

func MethodsDemo() {
	fmt.Println("-------MethodsDemo--------")
	mack := user{"mack", "mack@email.com"}
	mack.notify()
	mack.changeEmail("mack@hotmail.com")
	(&mack).changeEmail("mack@gmail.com")

	tom := &user{"tom", "tom@email.com"}
	tom.notify()
	(*tom).notify()
	tom.changeEmail("tom@hotmail.com")
	fmt.Println("---------------------------------")
}
