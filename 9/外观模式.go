package main

import "fmt"

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

type HomePlayerFacade struct {
	tv TV
	vb VoiceBox
	light Light
}

func (hp *HomePlayerFacade) DoKTV()  {
	fmt.Println("家庭影院进入KTV模式")
	hp.tv.On()
	hp.vb.On()
	hp.light.On()
}

func main() {
	homePlayer := new(HomePlayerFacade)

	homePlayer.DoKTV()
}