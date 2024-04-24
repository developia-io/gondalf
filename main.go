package main

import (
	"log"
)

type Recipe struct {
	ProjectName    string
	Url            string
	Install        bool
	PackageManager PackageManager
}

func main() {
	var recipe Recipe

	err := RunForm(&recipe)
	if err != nil {
		log.Fatal(err)
	}

	err = CreateProject(recipe)
	if err != nil {
		log.Fatal(err)
	}
}
