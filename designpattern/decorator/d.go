package main

import (
	"fmt"
	"math"
)

//装饰者模式，就是层层套娃处理
func main() {
	f := DecFunc(Double(Sqrt(nil)))
	fmt.Println(f(16.0))
}

type DecoratorFunc func(float64) float64

func DecFunc(dec DecoratorFunc) DecoratorFunc {
	return func(f float64) float64 {
		result := dec(f)
		return result
	}
}

func Double(decoratorFunc DecoratorFunc) DecoratorFunc {
	return func(f float64) float64 {
		var result float64 = f
		if decoratorFunc != nil {
			result = decoratorFunc(f)
		}
		return result * 2
	}
}

func Sqrt(decoratorFunc DecoratorFunc) DecoratorFunc {
	return func(f float64) float64 {
		var result float64 = f
		if decoratorFunc != nil {
			result = decoratorFunc(f)
		}
		return math.Sqrt(result)
	}
}
