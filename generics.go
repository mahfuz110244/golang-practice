package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

//  get error as string compare with error
// invalid operation: a < b (type parameter T is not comparable with <)
func minGenericInterface[T interface{}]( a T, b T) T {
	return b
}

func minGenericAny[T any]( a T, b T) T {
	return b
}

func minGenericFloat64[T ~float64]( a T, b T) T {
	if a < b {
		return a
	}
	return b
}
 
type minTypes interface {
	~ float64 | int | string
}
func min[T minTypes]( a T, b T) T {
	if a < b {
		return a
	}
	return b
}


func main() {
	fmt.Println(minInt(1, 2))
	fmt.Println(minFloat64(1.0, 2.0))

	fmt.Println(minGenericInterface(1, 2))
	fmt.Println(minGenericInterface(1.0, 2.0))

	fmt.Println(minGenericAny[int](1, 2))
	fmt.Println(minGenericAny[float64](1.0, 2.0))
	
	type superFloat float64
	var sf superFloat = 0.3
	fmt.Println(minGenericFloat64(sf, 0.2))

	fmt.Println(min(1, 2))
	fmt.Println(min("1.0", "2.0"))
}
