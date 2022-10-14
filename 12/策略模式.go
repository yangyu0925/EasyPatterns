package main

import "fmt"

//type WeaponStrategy interface {
//	UseWeapon()
//}
//
//type Ak47 struct {}
//
//func (ak *Ak47) UseWeapon()  {
//	fmt.Println("使用Ak47 去战斗")
//}
//
//type Knifr struct {}
//
//func (ak *Knifr) UseWeapon() {
//	fmt.Println("使用匕首 去战斗")
//}
//
//type Hero struct {
//	strategy WeaponStrategy
//}
//
//func (h *Hero) SetWeaponStgrategy(s WeaponStrategy)  {
//	h.strategy = s
//}
//
//func (h *Hero) Fight()  {
//	h.strategy.UseWeapon()
//}
//
//func main() {
//	hero := Hero{}
//
//	hero.SetWeaponStgrategy(new(Ak47))
//	hero.Fight()
//
//	hero.SetWeaponStgrategy(new(Knifr))
//	hero.Fight()
//}

type SellStrategy interface {
	GetPrice(price float64) float64
}

type StrategyA struct {}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A, 所有商品打八折")
	return price * 0.8
}

type StrategyB struct {}

func (sa *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("执行策略B, 所有商品满200 减100")

	if price >= 200 {
		price -= 100
	}

	return price
}

type Goods struct {
	Price float64
	Strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy)  {
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值 ", g.Price , " .")
	return g.Strategy.GetPrice(g.Price)
}

func main() {
	nike := Goods{
		Price: 200.00,
	}

	nike.SetStrategy(new(StrategyA))
	fmt.Println("上午", nike.SellPrice())

	nike.SetStrategy(new(StrategyB))
	fmt.Println("下午", nike.SellPrice())

}