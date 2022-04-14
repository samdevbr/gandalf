package main

import (
	"fmt"
	"github.com/samdevbr/gandalf/spec"
	"log"
)

func main() {
	specification, err := spec.New("gandalf.toml")

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(specification.PHP.Framework)
}
