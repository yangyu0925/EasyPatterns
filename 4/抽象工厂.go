package main

import "fmt"

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

type ChinaApple struct {

}

func (ca *ChinaApple) ShowApple()  {
	fmt.Println("中国苹果")
}

type ChinaBanana struct {

}

func (ca *ChinaBanana) ShowBanana()  {
	fmt.Println("中国香蕉")
}

type ChinaPear struct {

}

func (ca *ChinaPear) ShowPear()  {
	fmt.Println("中国梨")
}

type ChinaFactory struct {}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(ChinaApple)

	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(ChinaBanana)

	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(ChinaPear)

	return pear
}


type JapanApple struct {}

func (ja *JapanApple) ShowApple()  {
	fmt.Println("日本苹果")
}

type JapanBanana struct {}

func (ja *JapanBanana) ShowBanana()  {
	fmt.Println("日本香蕉")
}

type JapanPear struct {}

func (ja *JapanPear) ShowPear()  {
	fmt.Println("日本梨")
}

type JapanFactory struct {}

func (jf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple

	apple = new(JapanApple)

	return apple
}

func (jf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana

	banana = new(JapanBanana)

	return banana
}

func (jf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear

	pear = new(JapanPear)

	return pear
}

func main() {
	var JFac AbstractFactory
	JFac = new(JapanFactory)

	var JApple AbstractApple
	JApple = JFac.CreateApple()
	JApple.ShowApple()

	var JBanana AbstractBanana
	JBanana = JFac.CreateBanana()
	JBanana.ShowBanana()

	var JPear AbstractPear
	JPear = JFac.CreatePear()
	JPear.ShowPear()





	CFac := new(ChinaFactory)
	var CApple AbstractApple
	CApple = CFac.CreateApple()
	CApple.ShowApple()

	var CBanana AbstractBanana
	CBanana = CFac.CreateBanana()
	CBanana.ShowBanana()

	var CPear AbstractPear
	CPear = CFac.CreatePear()
	CPear.ShowPear()

}

