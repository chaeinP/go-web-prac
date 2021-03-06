package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	tpl, err := template.ParseFiles("../../data/tpl.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
