package copier

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	type Person struct {
		Name       string
		MiddleName *string
		Surname    string
		A          func()
	}

	type User struct {
		Person
		Email   string
		Age     int8
		Married bool
	}

	type Employee struct {
		Name       string
		MiddleName string
		Surname    string
		Email      string
		Age        *string
		A          string
		Married    string
	}

	src := User{
		Person: Person{
			Name:       "John",
			MiddleName: nil,
			Surname:    "Smith",
		},
		Email:   "john.smith@joy.me",
		Age:     33,
		Married: false,
	}
	dst := Employee{}

	copier := New(Tag("json"), Skip())
	copier.Copy(&dst, &src)

	fmt.Printf("%#v\n", dst)

	dst.Married = "ON"

	var a User
	copier.Copy(&a, &dst)

	fmt.Printf("%#v\n", &a)
}
