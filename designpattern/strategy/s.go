package main

import "fmt"

//策略模式，这其实就是工厂模式多了个判断
func main() {
	money := 100.0
	cc := NewCashContext("打八折")
	money = cc.GetMoney(money)
	fmt.Println("100打八折实际金额为", money)
	money = 199
	cc = NewCashContext("满一百返20")
	money = cc.GetMoney(money)
	fmt.Println("199满一百返20实际金额为", money)
	money = 199
	cc = NewCashContext("没有折扣")
	money = cc.GetMoney(money)
	fmt.Println("199没有折扣实际金额为", money)
}

type cashSuper interface {
	AcceptMoney(money float64) float64
}

type cashNormal struct{}

func newCashNormal() cashNormal {
	instance := new(cashNormal)
	return *instance
}

func (c cashNormal) AcceptMoney(money float64) float64 {
	return money
}

type cashRebate struct {
	Rebate float64 //折扣
}

func newCashRebate(rebate float64) cashRebate {
	instance := new(cashRebate)
	instance.Rebate = rebate
	return *instance
}

func (c cashRebate) AcceptMoney(money float64) float64 {
	return money * c.Rebate
}

type cashReturn struct {
	MoneyCondition float64
	MoneyReturn    float64
}

func newCashReturn(moneyCondition float64, moneyReturn float64) cashReturn {
	instance := new(cashReturn)
	instance.MoneyCondition = moneyCondition
	instance.MoneyReturn = moneyReturn
	return *instance
}

func (c cashReturn) AcceptMoney(money float64) float64 {
	if money >= c.MoneyCondition {
		moneyMinus := int(money / c.MoneyCondition)
		return money - float64(moneyMinus)*c.MoneyReturn
	}
	return money
}

type CashContext struct {
	Strategy cashSuper
}

func NewCashContext(cashType string) CashContext {
	c := new(CashContext)
	//这里事实上是简易工厂模式的变形，用来生产策略
	switch cashType {
	case "打八折":
		c.Strategy = newCashRebate(0.8)
	case "满一百返20":
		c.Strategy = newCashReturn(100.0, 20.0)
	default:
		c.Strategy = newCashNormal()
	}
	return *c
}

func (c CashContext) GetMoney(money float64) float64 {
	return c.Strategy.AcceptMoney(money)
}
