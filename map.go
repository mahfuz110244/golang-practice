/*func make(t Type, size ...IntegerType) Type
The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type:

Slice: The size specifies the length. The capacity of the slice is
equal to its length. A second integer argument may be provided to
specify a different capacity; it must be no smaller than the
length. For example, make([]int, 0, 10) allocates an underlying array
of size 10 and returns a slice of length 0 and capacity 10 that is
backed by this underlying array.

Map: An empty map is allocated with enough space to hold the
specified number of elements. The size may be omitted, in which case
a small starting size is allocated.

Channel: The channel's buffer is initialized with the specified
buffer capacity. If zero, or the size is omitted, the channel is
unbuffered.
*/

package main

import "fmt"

func main() {
	m := map[string]float64{"pi": 3.14}
	v, found := m["pi"] // v == 3.14  found == true
	fmt.Println(v, found)
	v, found = m["pie"] // v == 0.0   found == false
	fmt.Println(v, found)
	_, found = m["pi"] // found == true
	fmt.Println(found)

	if v, found := m["pie"]; found {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
	fmt.Println(m["pi"])
	fmt.Println(m["pie"])
	m["pie"] = 3.1416
	fmt.Println(m["pie"])
	m["pie"] = 3.14161455 // update key value
	fmt.Println(m["pie"])
	var m1 map[string]int // nil map of string-int pairs

	for key, value := range m { // Order not specified
		fmt.Println(key, value)
	}

	m2 := make(map[string]int)         // Empty map of string-float64 pairs
	m3 := make(map[string]string, 100) // Preallocate room for 100 entries
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}
