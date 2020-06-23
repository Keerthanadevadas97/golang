package main

import (
	"fmt"

	"./animal"
)

func main() {
	k := animal.Animal{}
	k.SetName("cat", "keerthana")
	k.GetName()
	fmt.Println(k.GetName())
	fmt.Println(k.GetOwner())

}
