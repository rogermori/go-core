package main

import (
	"Markdown/pages"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}
	fileName := os.Args[1]
	p, err := pages.NewPage(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	err = p.Render("layout.html", os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}

}