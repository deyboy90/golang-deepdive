package main

import (
	"github.com/deyboy90/golang-deepdive/dive4-composition/assertions"
	"github.com/deyboy90/golang-deepdive/dive4-composition/decoupling/style1"
	"github.com/deyboy90/golang-deepdive/dive4-composition/decoupling/style2"
	"github.com/deyboy90/golang-deepdive/dive4-composition/decoupling/style3"
	"github.com/deyboy90/golang-deepdive/dive4-composition/grouping/behavior"
	"github.com/deyboy90/golang-deepdive/dive4-composition/grouping/state"
)

func main() {
	behavior.GroupingByBehaviorDemo()
	state.GroupingByStateDemo()

	// Decoupling styles
	style1.Style1Demo()
	style2.Style2Demo()
	style3.Style3Demo()

	assertions.TypeAssertionsDemo()
}
