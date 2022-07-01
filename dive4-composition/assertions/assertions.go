package assertions

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

func (bike) Move() {
	fmt.Println("Moving the bike")
}

func (bike) Lock() {
	fmt.Println("Locking the bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func TypeAssertionsDemo() {
	var ml MoveLocker
	var m Mover

	ml = bike{}
	// we can assign the concrete type bike which implements MoveLocker interface
	// to Mover because any concrete type of MoveLocker interface implements all
	// the methods that are needed by Mover interface
	m = ml
	m.Move()

	// cannot use m (variable of type Mover) as MoveLocker value in assignment:
	// Mover does not implement MoveLocker (missing method Lock)compilerInvalidIfaceAssign
	// This does not work because the compiler cannot guarantee if the concrete value
	// store inside m knows how to Lock and Unlock
	// ml = m

	// if there is a bike value stored inside of m at the moment the code is executed.
	// If there is, then the variable b is given a copy of the bike value stored.
	// Then the copy can be copied inside of the ml interface variable.
	if b, ok := m.(bike); ok {
		ml = b
		ml.Move()
	}
}
