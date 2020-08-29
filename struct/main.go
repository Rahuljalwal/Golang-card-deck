package main

import "fmt"

type contactInfo struct {
	Email string
	zip   int
}

// declaring struct
type Person struct {
	FirstName string
	LastName  string
	contact   contactInfo
}

func main() {

	//var info Person

	info := Person{
		FirstName: "Rahul",
		LastName:  "Jangra",
		contact: contactInfo{
			Email: "gmail.com",
			zip:   126102}}

	infoPointer := &info
	infoPointer.update()
	info.print()

}

func (p Person) print() {
	fmt.Println(p)
}

func (p *Person) update() {
	(*p).FirstName = "new value"
}
