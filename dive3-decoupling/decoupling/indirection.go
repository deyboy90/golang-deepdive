package decoupling

import "fmt"

type data struct {
	name string
	age  int
}

// method defined using value semantics
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// method defined using pointer semantics
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func IndirectionDemo() {
	fmt.Println("-------IndirectionDemo--------")
	d := data{
		name: "Daniel",
		age:  24,
	}

	// since displayName method works on value semantics,
	// when it is assigned to f1 it oeprates on its own
	// copy of d struct, which means any changes to the
	// original data will not reflect on f1's own copy of d
	f1 := d.displayName
	f1()
	d.name = "Ramsey"
	f1()

	// here we will assign the setAge method to f2 which
	// works on pointer semantics, which means that f2
	// will point to the same d and will not create its
	// own copy, hence changes to d will be picked up
	// by f2 setAge method calls
	d = data{
		name: "Daniel",
		age:  24,
	}
	f2 := d.setAge
	f2(32)
	d.name = "Ramsey"
	f2(32)
	fmt.Println("---------------------------------")

}
