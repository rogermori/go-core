package main

import "fmt"

var (
	author = "Roger Mori"
	Version = "0.0.1"
)

const (
	CCVisa = "Visa"
	CCMasterCard = "MasterCard"
)

func main() {
	var greeting string = "Hello World!"
	var a, b, c int = 1, 2, 3
	fmt.Println(greeting, a, b, c)

	var dirName = "Roger Mori"
	var d, e, f = 1, 2.0, "Three"
	fmt.Println(dirName,d,e,f)

	course := "Essential Go"
	x, y , z := 1, 2, 3

	fmt.Println(course,x,y,z)

	fmt.Println(author,Version)

}
