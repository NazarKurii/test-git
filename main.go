package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) setName(name string) *Test {
	if t == nil {
		t = &Test{}
	}

}

func main() {
	var t *Test
	t.setName("name")
	fmt.Println(t)
}
