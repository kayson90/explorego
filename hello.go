package main

import "fmt"

type Men interface {
	Print()
	Set()
}

type Employee struct {
	company string
}

type Women struct {
}

func (e *Employee) Set() {
	e.company = "alibaba"
}

func (e *Employee) Print() {
	fmt.Printf("company = %v\n", e.company)
}

var e = Employee{
	company: "tencent",
}

func create() Men {
	return &e
}

func main() {
	fmt.Println("hello")
	o := create()
	e1, ok := o.(*Employee)
	if ok {
		e1.Set()
		e1.Print()
	}
	e.Print()
}
