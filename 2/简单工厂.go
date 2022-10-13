package main

import "fmt"

type Fruit interface {
	Show()
}

type Apple struct {

}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {

}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {

}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

type Factory struct {

}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
