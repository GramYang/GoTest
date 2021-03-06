package main

import "fmt"

//状态模式，子类传入不同的实现，父类的方法调用子类的方法，因而也会有不同的实现
func main() {
	tWork := Work{}
	tState := ForenoonState{}
	tWork.SetState(tState)
	tWork.SetFinishState(true)
	tWork.SetHour(22)
	tWork.WriteProgram()
}

type Work struct {
	hour    int
	current State
	finish  bool
}

func (w *Work) SetState(s State) {
	w.current = s
}

func (w *Work) SetHour(hour int) {
	w.hour = hour
}

func (w *Work) SetFinishState(finish bool) {
	w.finish = finish
}

func (w Work) WriteProgram() {
	w.current.WriteProgram(w)
}

type State interface {
	WriteProgram(work Work)
}

type ForenoonState struct{}

func (fs ForenoonState) WriteProgram(work Work) {
	if work.hour < 12 {
		fmt.Println("上午")
	} else {
		work.SetState(NoonState{})
		work.WriteProgram()
	}
}

type NoonState struct{}

func (ns NoonState) WriteProgram(work Work) {
	if work.hour < 13 {
		fmt.Println("中午")
	} else {
		work.SetState(AfternoonState{})
		work.WriteProgram()
	}
}

type AfternoonState struct{}

func (as AfternoonState) WriteProgram(work Work) {
	if work.hour < 17 {
		fmt.Println("下午")
	} else {
		work.SetState(EveningState{})
		work.WriteProgram()
	}
}

type EveningState struct{}

func (es EveningState) WriteProgram(work Work) {
	if work.finish {
		work.SetState(RestState{})
		work.WriteProgram()
	} else {
		if work.hour < 21 {
			fmt.Println("晚间")
		} else {
			work.SetState(SleepingState{})
			work.WriteProgram()
		}
	}
}

type SleepingState struct{}

func (ss SleepingState) WriteProgram(work Work) {
	fmt.Println("睡着了")
}

type RestState struct{}

func (rs RestState) WriteProgram(work Work) {
	fmt.Println("下班回家")
}
