package main

import "fmt"

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Printf("%v prrrrr\n",c.name)
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Printf("%v woof woof\n",d.name)
}

func (d Dog) Name() string {
	return d.name
}

type Animal interface {
	Pet()
	Name() string
}

func Compliment(a Animal) {
	fmt.Println("Great Job", a.Name())
	a.Pet()
}


func main() {
	c := Cat{"johnny larry"}
	Compliment(c)
	
	d := Dog{"Trixie Belle"}
	Compliment(d)
	
}