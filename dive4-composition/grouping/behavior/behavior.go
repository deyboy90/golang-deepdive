// -----------------
// Grouping By Behavior
// -----------------

package behavior

import "fmt"

// When an interface has one method then name it as a verb
type Speakers interface {
	Speak()
}

// Dog contains everything from Animal, plus specific attributes that only a Dog has.
type Dog struct {
	Name       string // duplicated
	IsMammal   bool   // duplicated
	PackFactor int
}

// Speak knows how to speak like a dog.
func (d *Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything from Animal, plus specific attributes that only a Cat has.
type Cat struct {
	Name        string // duplicated
	IsMammal    bool   // duplicated
	ClimbFactor int
}

// Speak knows how to speak like a cat.
func (c *Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func Execute() {
	// Create a Dog by initializing its Animal parts and then its specific Dog attributes.
	dog := &Dog{
		Name:       "Fido",
		IsMammal:   true,
		PackFactor: 5,
	}

	// Create a Cat by initializing its Animal parts and then its specific Cat attributes.
	cat := &Cat{
		Name:        "Milo",
		IsMammal:    true,
		ClimbFactor: 4,
	}

	speakers := []Speakers{dog, cat}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}

// ----------
// Conclusion
// ----------

// This code smells bad because:
// - The Animal type provides an abstraction layer of reusable state.
// - The program never needs to create or solely use a value of Animal type.
// - The implementation of the Move method for the Animal type is generalization.
