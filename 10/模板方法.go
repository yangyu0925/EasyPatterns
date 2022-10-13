package main

import "fmt"

type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()

	WatAddThings() bool
}

type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.WatAddThings() == true {
		t.b.AddThings()
	}
}

type MakeCaffee struct {
	template
}

func NewMakeCaffee() *MakeCaffee {
	makeCaffe := new(MakeCaffee)
	makeCaffe.b = makeCaffe
	return makeCaffe
}

func (mc *MakeCaffee) BoilWater()  {
	fmt.Println("将水煮到100摄氏度")
}

func (mc *MakeCaffee) Brew()  {
	fmt.Println("用水冲咖啡豆")
}

func (mc *MakeCaffee) PourInCup()  {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

func (mc *MakeCaffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (mc *MakeCaffee) WatAddThings() bool {
	return true
}

type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func (mt *MakeTea) BoilWater()  {
	fmt.Println("将水煮到80摄氏度")
}

func (mt *MakeTea) Brew()  {
	fmt.Println("用水冲茶叶")
}

func (mt *MakeTea) PourInCup()  {
	fmt.Println("将充好的咖啡倒入茶壶中")
}

func (mt *MakeTea) AddThings()  {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) WatAddThings() bool {
	return false
}

func main() {
	makeCoffee := NewMakeCaffee()
	makeCoffee.MakeBeverage()

	fmt.Println("------------")
	MakeTea := NewMakeTea()
	MakeTea.MakeBeverage()
}