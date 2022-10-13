package main

import "fmt"

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

type Apple struct {
	Fruit
}

func (apple *Apple) Show()  {
	fmt.Println("我是苹果")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show()  {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

type JapanApple struct {
	Fruit
}

func (jp *JapanApple) Show() {
	fmt.Println("我是日本苹果")
}

type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	fruit = new(Apple)

	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit

	fruit = new(Banana)

	return fruit
}

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit

	fruit = new(Pear)

	return fruit
}

type JapanAppleFactory struct {
	AbstractFactory
}

func (fac *JapanAppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	fruit = new(JapanApple)

	return fruit
}

func main() {
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)

	var apple Fruit
	apple = appleFac.CreateFruit()
	apple.Show()

	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)

	var banana Fruit
	banana = bananaFac.CreateFruit()
	banana.Show()

	var pearFac AbstractFactory
	pearFac = new(PearFactory)

	var pear Fruit
	pear = pearFac.CreateFruit()
	pear.Show()

	var JapanAppleFac AbstractFactory
	JapanAppleFac = new(JapanAppleFactory)

	var japanApple Fruit
	japanApple = JapanAppleFac.CreateFruit()
	japanApple.Show()

}



