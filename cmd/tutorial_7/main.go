package main

import "fmt"

func updateName(s *string) string {
	*s = "NA"
	return *s
}

func main() {

	name := "MP"
	m := &name
	fmt.Println(updateName(m))
	fmt.Println(name)
}
