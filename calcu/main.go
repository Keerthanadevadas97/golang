package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	var option string
	res := 0
	fmt.Scanln(&n1, &n2)
	fmt.Println("Enter your option:\n 1.add\n 2.sub\n 3.mul\n 4.div")
	fmt.Scanln(&option)
	switch option {
	case "1":
		res = n1 + n2
		break
	case "2":
		res = n1 - n2
		break
	case "3":
		res = n1 * n2
		break
	case "4":
		res = n1 / n2
		break
	default:
		fmt.Println("invalid option")
	}
	fmt.Println("result=", res)
}
