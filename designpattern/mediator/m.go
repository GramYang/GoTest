package main

import "fmt"

//中介者模式，用一个中介对象来封装一系列的对象交互。
//中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间地交互。
func main() {
	//创建一个具体中介者
	tMediator := UnitedNationsSecurityCouncil{}
	//创建具体同事,并且让他认识中介者
	tColleageA := USA{
		UnitedNations: tMediator,
	}
	tColleageB := Iraq{
		UnitedNations: tMediator,
	}
	//让中介者认识每一个具体同事
	tMediator.USA = tColleageA
	tMediator.Iraq = tColleageB
	//A同事发送消息
	tColleageA.SendMessage("停止核武器研发，否则发动战争")  //伊拉克收到对方消息: 停止核武器研发，否则发动战争
	tColleageB.SendMessage("我们没有研发核武器，也不怕战争") //美国收到对方消息: 我们没有研发核武器，也不怕战争
}

type UnitedNations interface {
	ForwardMessage(message string, country Country)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Iraq
}

func (unsc UnitedNationsSecurityCouncil) ForwardMessage(message string, country Country) {
	switch country.(type) {
	case USA:
		unsc.Iraq.GetMessage(message)
	case Iraq:
		unsc.USA.GetMessage(message)
	default:
		fmt.Println("The country is not a member of UNSC")
	}
}

type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}
type USA struct {
	UnitedNations
}

func (usa USA) SendMessage(message string) {
	usa.UnitedNations.ForwardMessage(message, usa)
}

func (usa USA) GetMessage(message string) {
	fmt.Printf("美国收到对方消息: %s\n", message)
}

type Iraq struct {
	UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
	iraq.UnitedNations.ForwardMessage(message, iraq)
}

func (iraq Iraq) GetMessage(message string) {
	fmt.Printf("伊拉克收到对方消息: %s\n", message)
}
