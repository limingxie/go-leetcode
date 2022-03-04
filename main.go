package main

import "fmt"

func main() {
	var s string = ""
	s = "()"
	a := IsValid(s)
	fmt.Println(a)

	s = "()[]{}"
	a1 := IsValid(s)
	fmt.Println(a1)

	s = "(]"
	a2 := IsValid(s)
	fmt.Println(a2)

	s = "([)]"
	a3 := IsValid(s)
	fmt.Println(a3)

	s = "{[]}"
	a4 := IsValid(s)
	fmt.Println(a4)

	s = "}"
	a5 := IsValid(s)
	fmt.Println(a5)

	s = "(){}}{"
	a6 := IsValid(s)
	fmt.Println(a6)

}
