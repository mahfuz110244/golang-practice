// https://www.geeksforgeeks.org/inheritance-in-golang/#:~:text=Since%20Golang%20does%20not%20support,No%20Inheritance%20Concept%20in%20Golang.
// https://www.youtube.com/watch?v=I0Ei70cZtf8&ab_channel=GolangDojo

package main

import "fmt"

type first struct {
	baseOne string
}

type second struct {
	baseTwo string
}

func (f first) printBaseOne() string {
	return f.baseOne
}

func (s second) printBaseTwo() string {
	return s.baseTwo
}

type childOne struct {
	first
}

type childTwo struct {
	second
}

type child struct {
	first
	second
}

func main() {
	c1 := childOne{
		first{
			baseOne: "In base one struct for child one",
		},
	}

	c2 := childTwo{
		second{
			baseTwo: "In base two struct for child two",
		},
	}

	c := child{
		first{
			baseOne: "In base one struct for child",
		},
		second{
			baseTwo: "In base two struct for child",
		},
	}

	fmt.Println(c1.printBaseOne())
	fmt.Println(c2.printBaseTwo())
	fmt.Println(c.printBaseOne())
	fmt.Println(c.printBaseTwo())

}
