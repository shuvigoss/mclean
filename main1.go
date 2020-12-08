package main

import "fmt"

func main() {
	var a = make([]string, 0)
	a = append(a, "1")
	a = append(a, "2")
	a = append(a, "3")
	a = append(a, "4")

	slice(&a)

	fmt.Println("a")
}

func slice(s *[]string) {
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == "1" {
			*s = append((*s)[:i], (*s)[i+1:]...)
			i-- // form the remove item index to start iterate next item
		}
	}

}
