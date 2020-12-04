package main

import "fmt"

//享元模式，一个结构体用map存储子结构体+状态模式
func main() {
	//  创建享元
	fly := NewFlyweight()
	base := NewPeopleBase()
	fly.SetElement("PeopleBase", base)
	//  生成两个人物
	people_1 := fly.GetElement("PeopleBase").(IProperty)
	people_2 := fly.GetElement("PeopleBase").(IProperty)
	//  捡到装备了——装饰者
	people_1 = NewHelmet(people_1, 10, 10)
	people_2 = NewHelmet(people_2, 100, 100)
	//  获取它们血量上限和魔法上限
	hp_1 := people_1.GetHPLimit()
	mp_1 := people_1.GetMPLimit()
	hp_2 := people_2.GetHPLimit()
	mp_2 := people_2.GetMPLimit()
	fmt.Printf("People_1:\nHP:%d\nMP:%d\n", hp_1, mp_1)
	fmt.Printf("People_2:\nHP:%d\nMP:%d\n", hp_2, mp_2)
}

type IProperty interface {
	GetHPLimit() int
	GetMPLimit() int
}

type PeopleBase struct {
	MAX_HP int
	MAX_MP int
}

func NewPeopleBase() *PeopleBase {
	return &PeopleBase{100, 100}
}

func (this *PeopleBase) GetHPLimit() int {
	return this.MAX_HP
}

func (this *PeopleBase) GetMPLimit() int {
	return this.MAX_MP
}

type Helmet struct {
	base   IProperty
	HP_ADD int
	MP_ADD int
}

func (this *Helmet) GetHPLimit() int {
	return this.base.GetHPLimit() + this.HP_ADD
}

func (this *Helmet) GetMPLimit() int {
	return this.base.GetMPLimit() + this.MP_ADD
}

func NewHelmet(property IProperty, hp_add, mp_add int) *Helmet {
	return &Helmet{property, hp_add, mp_add}
}

type Element struct {
	Value interface{}
}

func newElement(value interface{}) *Element {
	return &Element{value}
}

type FlyweightFactory struct {
	pool map[string]*Element
}

func (this *FlyweightFactory) GetElement(key string) interface{} {
	return this.pool[key].Value
}

func (this *FlyweightFactory) SetElement(key string, value interface{}) {
	ne := newElement(value)
	this.pool[key] = ne
}

func NewFlyweight() *FlyweightFactory {
	flyweight := FlyweightFactory{}
	flyweight.pool = make(map[string]*Element)
	return &flyweight
}
