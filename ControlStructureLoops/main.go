package main

import "fmt"

func main() {

	if true == true {
		fmt.Println("true is true")
	} else {
		fmt.Println("true is false")
	}

	if true == false {
		fmt.Println("true is false")
	} else {
		fmt.Println("true is true")
	}

    //	for-loop C style
    for i:= 0; i < 5; i++ {
		fmt.Println(i)
	}

    for {
		fmt.Println("Go!")
		break
	}

	j := true
    for j {
		fmt.Println("This should happen only once")
		j = false
	}

}