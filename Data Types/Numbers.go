package main

import "fmt"

func main() {
	var value int
	fmt.Println("1.int\n2.uint\n3.float")
	fmt.Scan(&value)
	switch value {
	case 1: // Can contain signed and unsigned integers
		var a int
		var a1 int8
		var a2 int16
		var a3 int32 //rune is same as int32
		var b3 rune
		var a4 int64
		// max value each type can print
		a = 10
		a1 = 100
		a2 = -10000
		a3 = 1000000000
		b3 = 1000000000
		if a3 == b3 {
			fmt.Println("int32 and rune are same")
		}
		a4 = -1000000000000000000
		fmt.Println(" a:", a, "\n", "a1:", a1, "\n", "a2:", a2, "\n", "a3:", a3, "\n", "a4:", a4)
	case 2: // Contains only signed integers
		var a uint
		var a1 uint8
		var b1 byte
		var a2 uint16
		var a3 uint32
		var a4 uint64
		a = 10
		a1 = 100
		b1 = 100
		if b1 == a1 {
			fmt.Println("uint8 and byte are same")
		}
		a2 = 10000
		a3 = 1000000000
		a4 = 1000000000000000000
		fmt.Println(" a:", a, "\n", "a1:", a1, "\n", "a2:", a2, "\n", "a3:", a3, "\n", "a4:", a4)
	case 3:
		var a float32
		var a1 float64
		var a2 complex64
		// var a3 complex128 this can occupy quite bigger
		a = 1.01 - 0.99
		a1 = 1.01 - 0.99
		a2 = 100000000000000000000000000000000000000 + 100000000000000000000000000000000000000i
		fmt.Println("a: ", a, "\n", "a1: ", a1, "\n", "a2 :", a2)

	}

}
