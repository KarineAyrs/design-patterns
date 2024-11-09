package main

import "fmt"

// Animal is the type for abstract factory.
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is the concrete factory for dogs.
type Dog struct {
}

func (d *Dog) Says() {
	fmt.Println("Woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

// Cat is the concrete factory for cats.
type Cat struct {
}

func (c *Cat) Says() {
	fmt.Println("Meow")
}

func (c *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("a dog likes water: ", dog.LikesWater())
	fmt.Println("a cat likes water: ", cat.LikesWater())
}
