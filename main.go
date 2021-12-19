package main

import (
	"fmt"

	"github.com/akedev7/GoWithCustomType/organization"
)

func main() {
	//Custom data type
	p := organization.NewPerson("James", "Wilson")
	println(p.ID())
	println(p.FullName())
	err := p.SetTwitterHandler("jame wil")
	if err != nil {
		fmt.Printf("An error occur setting twitter handler %s \n", err.Error())
	}

}
