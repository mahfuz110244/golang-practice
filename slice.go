// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	/*For example, make([]int, 2, 10) allocates an underlying array
	of size 10 and returns a slice of length 2 and capacity 10 that is
	backed by this underlying array.*/
	a := make([]int, 2, 2)
	fmt.Println(a)

	a[0] = 11
	a[1] = 12
	fmt.Println(a)

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5].

	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2].

	l = s[2:]
	fmt.Println("sl3:", l)

	//We can declare and initialize a variable for slice in a single line as well.

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

}
