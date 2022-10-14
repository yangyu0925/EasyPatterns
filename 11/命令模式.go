package main

import (
	"fmt"
)

//type Doctor struct {}
//
//func (d *Doctor) treatEye() {
//	fmt.Println("医生治疗眼睛")
//}
//
//func (d *Doctor) treatNose() {
//	fmt.Println("医生治疗鼻子")
//}
//
//type Command interface {
//	Treat()
//}
//
//type CommandTreateEye struct {
//	docker *Doctor
//}
//
//func (cmd *CommandTreateEye) Treat()  {
//	cmd.docker.treatEye()
//}
//
//type CommandTreatNose struct {
//	docter *Doctor
//}
//
//func (cmd *CommandTreatNose) Treat() {
//	cmd.docter.treatNose()
//}
//
//type Nurse struct {
//	CmdList []Command
//}
//
//func (n (Nurse)) Notify()  {
//	if n.CmdList == nil {
//		return
//	}
//
//	for _, cmd := range n.CmdList {
//		cmd.Treat()
//	}
//}
//
//func main() {
//	doctor := new(Doctor)
//	cmdEye := CommandTreateEye{
//		doctor,
//	}
//	cmdNose := CommandTreatNose{
//		doctor,
//	}
//
//	nurse := new(Nurse)
//
//	nurse.CmdList = append(nurse.CmdList, &cmdEye)
//	nurse.CmdList = append(nurse.CmdList, &cmdNose)
//
//	nurse.Notify()
//}

type Cooker struct {}

func (c *Cooker) MakeChinken()  {
	fmt.Println("烤串师傅烤了鸡肉串儿")
}

func (c *Cooker) MakeChuaner()  {
	fmt.Println("烤串师傅烤了羊肉串儿")
}

type Command interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make()  {
	cmd.cooker.MakeChinken()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

type WaiterMM struct {
	CmdList []Command
}

func (w *WaiterMM) Notify()  {
	if w.CmdList == nil {
		return
	}

	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker}
	cmdChuaner := CommandCookChuaner{cooker}

	mm := new(WaiterMM)
	mm.CmdList = append(mm.CmdList, &cmdChuaner)
	mm.CmdList = append(mm.CmdList, &cmdChicken)

	mm.Notify()
}