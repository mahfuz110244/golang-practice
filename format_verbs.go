package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// Integer
	num := 65

	// Float
	pi := 3.1415926535

	// String
	str := "Hello Go"

	// Rune
	ch := 'A'

	// Boolean
	ok := true

	// Pointer
	ptr := &num

	// Struct
	user := User{
		Name: "Mahfuz",
		Age:  30,
	}

	fmt.Println("========== Integer ==========")
	fmt.Printf("%%d  (Decimal)           : %d\n", num)
	fmt.Printf("%%b  (Binary)            : %b\n", num)
	fmt.Printf("%%08b (8-bit Binary)     : %08b\n", num)
	fmt.Printf("%%o  (Octal)             : %o\n", num)
	fmt.Printf("%%x  (Hex Lowercase)     : %x\n", num)
	fmt.Printf("%%X  (Hex Uppercase)     : %X\n", num)

	fmt.Println("\n========== Rune / Character ==========")
	fmt.Printf("%%c  (Character)         : %c\n", ch)
	fmt.Printf("%%q  (Quoted Character)  : %q\n", ch)
	fmt.Printf("%%U  (Unicode)           : %U\n", ch)

	fmt.Println("\n========== String ==========")
	fmt.Printf("%%s  (String)            : %s\n", str)
	fmt.Printf("%%q  (Quoted String)     : %q\n", str)

	fmt.Println("\n========== Float ==========")
	fmt.Printf("%%f  (Default Float)     : %f\n", pi)
	fmt.Printf("%%.2f (2 Decimal Places) : %.2f\n", pi)
	fmt.Printf("%%.4f (4 Decimal Places) : %.4f\n", pi)
	fmt.Printf("%%e  (Scientific)        : %e\n", pi)
	fmt.Printf("%%E  (Scientific Upper)  : %E\n", pi)

	fmt.Println("\n========== Boolean ==========")
	fmt.Printf("%%t                     : %t\n", ok)

	fmt.Println("\n========== Pointer ==========")
	fmt.Printf("%%p                     : %p\n", ptr)

	fmt.Println("\n========== Type ==========")
	fmt.Printf("%%T (num)               : %T\n", num)
	fmt.Printf("%%T (str)               : %T\n", str)
	fmt.Printf("%%T (user)              : %T\n", user)

	fmt.Println("\n========== Default ==========")
	fmt.Printf("%%v                     : %v\n", user)
	fmt.Printf("%%+v                    : %+v\n", user)
	fmt.Printf("%%#v                    : %#v\n", user)

	fmt.Println("\n========== Width / Padding ==========")
	fmt.Printf("|%5d|\n", 12)
	fmt.Printf("|%-5d|\n", 12)
	fmt.Printf("|%05d|\n", 12)

	fmt.Println("\n========== Percentage ==========")
	fmt.Printf("Success: 95%%\n")
}
