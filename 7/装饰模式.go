package main

import "fmt"

type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

type HuaWei struct {}

func (hw *HuaWei) Show()  {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct {}

func (xm *XiaoMi) Show()  {
	fmt.Println("秀出了XiaoMi手机")
}

type MoDecorator struct {
	Decorator
}

func (md *MoDecorator) Show()  {
	md.phone.Show()
	fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show()  {
	kd.phone.Show()
	fmt.Println("手机壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

func main() {
	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show()

	fmt.Println("-------")
	moHuaWei := NewMoDecorator(huawei)
	moHuaWei.Show()

	keHuawei := NewKeDecorator(huawei)
	keHuawei.Show()

	kemoHuawei := NewMoDecorator(keHuawei)
	kemoHuawei.Show()
}