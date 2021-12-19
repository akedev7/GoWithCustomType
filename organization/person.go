package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

//Type declaration ( it copies the fields of a type over to a new type, so you can extend the method for yourself)(can refer to a struct as well)
type socialSecurityNumber string

//it will implement interface implcitly for you from Identifiable (if there are the same 2 methods, it will error during the conpliation)
func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func NewSocialSecurityNumber(value string) Identifiable {
	return socialSecurityNumber(value)
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

//Type declaration ( it copies the fields of a type over to a new type, so you can extend the method for yourself)(can refer to a struct as well)
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Employee struct {
	Name
}

type Person struct {
	Name                          // Embbed struct (it's similar to "implement" in java but only fields, so the field become as a part of the object)
	twitterHandler TwitterHandler //Composition
	Identifiable                  // Embeded interface (it's similar to "implement" in java but only methods, so the method become as a part of the object)
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name:         Name{first: firstName, last: lastName},
		Identifiable: identifiable,
	}
}

func (p *Person) ID() string {

	return fmt.Sprintf("Person's Identifier: %s", p.Identifiable.ID())
}

// You need to make the pointer receiver to modify the object state ( you can have pointer for every method for consistency)
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
