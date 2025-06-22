package main

import (
	"fmt"
	"strings"
	"time"
)

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

func add() {}

func mihua() {

}

func mihuadknvd() {

}

func loh() {}

func loh2() {

}

<<<<<<< HEAD
type dateOfBirth struct {
	time.Time
}

func (dob *dateOfBirth) UnmarshalJSON(b []byte) error {
	// b is a JSON string literal, e.g. `"1990-05-15"`
	s := strings.Trim(string(b), `"`) // remove quotes

	// parse using the correct layout "2006-01-02"
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	dob.Time = t
	return nil
=======
func loh3() {

}

func different() {

>>>>>>> 87438fa (More commit 3)
}
