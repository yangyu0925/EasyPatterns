package main

import "fmt"

//type SubSystemA struct{}
//
//func (sa *SubSystemA) MethodA()  {
//	fmt.Println("子系统方法A")
//}
//
//type SubSysTemB struct {}
//
//func (sb *SubSysTemB) MethodB()  {
//	fmt.Println("子系统方法B")
//}
//
//type SubSysTemC struct {}
//
//func (sc *SubSysTemC) MethodC()  {
//	fmt.Println("子系统方法C")
//}
//
//type SubSysTemD struct {}
//
//func (sd *SubSysTemD) MethodD()  {
//	fmt.Println("子系统方法D")
//}
//
//type Facade struct {
//	a *SubSystemA
//	b *SubSysTemB
//	c *SubSysTemC
//	d *SubSysTemD
//}
//
//func (f *Facade) MethodOne()  {
//	f.a.MethodA()
//	f.b.MethodB()
//}
//
//func (f *Facade) MethodTwo()  {
//	f.c.MethodC()
//	f.d.MethodD()
//}
//
//func main() {
//	sa := new(SubSystemA)
//	sa.MethodA()
//	sb := new(SubSysTemB)
//	sb.MethodB()
//
//	fmt.Println("--------------")
//
//	f := Facade {
//		a: new(SubSystemA),
//		b: new(SubSysTemB),
//		c: new(SubSysTemC),
//		d: new(SubSysTemD),
//	}
//
//	f.MethodTwo()
//}

type TV struct {}

func (t *TV) On()  {
	fmt.Println("打开 电视机")
}

func (t *TV) Off()  {
	fmt.Println("关闭 电视机")
}

type VoiceBox struct {}

func (v *VoiceBox) On()  {
	fmt.Println("打开 音响")
}

func (v *VoiceBox) Off()  {
	fmt.Println("关闭 音响")
}

type Light struct {}

func (l *Light) On()  {
	fmt.Println("打开 灯光")
}

func (l *Light) Off()  {
	fmt.Println("关闭 灯光")
}

type Xbox struct {}

func (x *Xbox) On()  {
	fmt.Println("打开 游戏机")
}

func (x *Xbox) Off()  {
	fmt.Println("关闭 游戏机")
}

//麦克风
type MicroPhone struct {}

func (m *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}

func (m *MicroPhone) Off() {
	fmt.Println("关闭 麦克风")
}

//投影仪
type Projector struct {}

func (p *Projector) On() {
	fmt.Println("打开 投影仪")
}

func (p *Projector) Off() {
	fmt.Println("关闭 投影仪")
}

type HomePlayerFacade struct {
	tv TV
	vb VoiceBox
	light Light
	xbox Xbox
	mp MicroPhone
	pro Projector
}

func (hp *HomePlayerFacade) DoKTV()  {
	fmt.Println("家庭影院进入KTV模式")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.light.Off()
	hp.vb.On()
}

func (hp *HomePlayerFacade) DoGame()  {
	fmt.Println("家庭影院进入Game模式")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}

func main() {
	homePlayer := new(HomePlayerFacade)

	homePlayer.DoKTV()

	fmt.Println("------------")

	homePlayer.DoGame()
}