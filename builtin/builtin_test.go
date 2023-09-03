package builtin

import (
	"fmt"
)

func ExampleMin() {
	a := min(1, 2, 3)
	fmt.Println("min(1,2,3)=", a)

	b := min(1.0, 2.0, 3.0)
	fmt.Println("min(1.0,2.0,3.0)=", b)

	c := min(1, 2.0, 3)
	fmt.Println("min(1,2.0,3)=", c)

	d := min("a", "b", "c")
	fmt.Println(`min("a","b","c")=`, d)

	// _ = min(s...), this syntax is invalid, See: slices.Min

	// Output:
	// min(1,2,3)= 1
	// min(1.0,2.0,3.0)= 1
	// min(1,2.0,3)= 1
	// min("a","b","c")= a
}

func ExampleMax() {
	a := max(1, 2, 3)
	fmt.Println("min(1,2,3)=", a)

	b := max(1.0, 2.0, 3.0)
	fmt.Println("min(1.0,2.0,3.0)=", b)

	c := max(1, 2.0, 3)
	fmt.Println("min(1,2.0,3)=", c)

	d := max("a", "b", "c")
	fmt.Println(`min("a","b","c")=`, d)

	// _ = max(s...), this syntax is invalid, See: slices.Max

	// Output:
	// min(1,2,3)= 3
	// min(1.0,2.0,3.0)= 3
	// min(1,2.0,3)= 3
	// min("a","b","c")= c
}
