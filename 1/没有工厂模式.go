package main

import "fmt"

type Fruit struct {

}

func (f *Fruit) Show(name string)  {
	if name == "apple" {
		fmt.Println("我是苹果")
	} else if name == "banana" {
		fmt.Println("我是香蕉")
	} else if name == "pear" {
		fmt.Println("我是梨")
	}
}

func NewFruit(name string) *Fruit {
	fruit := new(Fruit)

	if name == "apple" {

	} else if name == "banana" {

	} else if name == "pear" {

	}

	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.Show("apple")

	banana := NewFruit("banana")
	banana.Show("banana")

	pear := NewFruit("pear")
	pear.Show("pear")
}
