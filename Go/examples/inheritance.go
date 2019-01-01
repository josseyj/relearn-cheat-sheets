package main

import "fmt"

type Named interface {
	getName() string
}

type Namable interface {
	setName(name string)
}

type AbstractItem interface {
	Named
	Namable
}

type ConcreteItem struct {
	name string
}

type Person struct {
	age int
	ConcreteItem
}

func (p *ConcreteItem) getName() string {
	return p.name
}

func (p *ConcreteItem) setName(name string) {
	p.name = name
}

func main() {
	var p = Person{age: 30}
	p.name = "Albert"
	processNamed(&p)
	processNamable(&p)
	processNamed(&p)

}

func processNamed(item Named) {
	fmt.Printf("Hello %s!\n", item.getName())
}

func processNamable(a Namable) {
	a.setName("Einstein")
}
