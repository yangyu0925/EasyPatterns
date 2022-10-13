package main

import "fmt"

type SellStrategy interface {
	GetPrice(price float64) float64
}

type StrategyA struct {}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A, 所有商品打八折")

	return price * 0.8
}

type StrategyB struct {}

func (sb *StrategyB) GetPrice(price float64) float64 {
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

func (g *Goods) SetStrateGy(s SellStrategy)  {
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值 ", g.Price , " .")
	return g.Strategy.GetPrice(g.Price)
}

func main() {
	nike := Goods{
		Price: 200.0,
	}

	nike.SetStrateGy(new(StrategyA))
	fmt.Println("上午nike鞋卖", nike.SellPrice())

	nike.SetStrateGy(new(StrategyB))
	fmt.Println("下午nike鞋卖", nike.SellPrice())

}
