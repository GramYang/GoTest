package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

func main() {
	//bit操作，目前golang的std并不支持直接解析和加密补码，需要你自己实现
	//b1()
	//10进制8进制转换
	//b2()
	//位运算方法测试
	//b3()
	//补码解析测试，这里限定2字节
	//b4()
	//补充一个go位运算符的测试
	b5()
}

func b1() {
	//10进制和2进制转换，用字符串功能有限
	b := "0011000111010101"
	a, err := strconv.ParseInt(b, 2, 16)
	fmt.Println(a, err) //12757 <nil>
	a1 := strconv.FormatInt(a, 2)
	fmt.Println(a1) //11000111010101
	a2, err := strconv.ParseInt(a1, 2, 16)
	fmt.Println(a2, err) //12757 <nil>
	//不能生成补码
	fmt.Println(strconv.FormatInt(-12757, 2)) //-11000111010101
	fmt.Println(fmt.Sprintf("%.16b", -12757)) //-0011000111010101
}

func b2() {
	a1 := fmt.Sprintf("%o", 18864)
	fmt.Println(a1) //44660
	a2, err := strconv.ParseInt(a1, 8, 16)
	fmt.Println(a2, err) //18864 <nil>
	a3 := fmt.Sprintf("%o", -18864)
	fmt.Println(a3) //-44660
	a4, err := strconv.ParseInt(a3, 8, 16)
	fmt.Println(a4, err) //-18864 <nil>
}

func b3() {
	//某一位取反
	var a uint8 = 20
	fmt.Println(strconv.FormatInt(int64(a), 2)) //00010100
	a1 := a ^ (1 << 4)
	fmt.Println(a1, strconv.FormatInt(int64(a1), 2)) //4 100
	a2 := a ^ (1 << 2)
	fmt.Println(a2, strconv.FormatInt(int64(a2), 2)) //16 10000
	a3 := a1 ^ (1 << 4)
	fmt.Println(a3, strconv.FormatInt(int64(a3), 2)) //20 10100
	a4 := a2 ^ (1 << 2)
	fmt.Println(a4, strconv.FormatInt(int64(a4), 2)) //20 10100
	//多位取反，比如取反1，3，5位
	b1 := []int{1, 3, 5}
	var b2 uint8 = 20
	for _, v := range b1 {
		b2 = b2 ^ (1 << v)
	}
	fmt.Println(b2, strconv.FormatInt(int64(b2), 2)) //62 00111110
	//某一位取0，&^就是位清空
	a5 := a &^ (1 << 4)
	fmt.Println(a5, strconv.FormatInt(int64(a5), 2)) //4 100
	a6 := a &^ (1 << 3)
	fmt.Println(a6) //20
	//多位置0，比如1，3，4
	b3 := []int{1, 3, 4}
	var b4 uint8 = 20
	for _, v := range b3 {
		b4 = b4 &^ (1 << v)
	}
	fmt.Println(b4, strconv.FormatInt(int64(b4), 2)) //4 100
	//某一位置1
	a7 := a | (1 << 4)
	fmt.Println(a7) //20
	a8 := a | (1 << 3)
	fmt.Println(a8, strconv.FormatInt(int64(a8), 2)) //28 11100
	//多位置1，比如1，2，3
	b5 := []int{1, 2, 3}
	var b6 uint8 = 20
	for _, v := range b5 {
		b6 = b6 | (1 << v)
	}
	fmt.Println(b6, strconv.FormatInt(int64(b6), 2)) //30 11110
	//取某一位的值
	a9 := a >> 3 & 1
	fmt.Println(a9) //0
	a10 := a >> 2 & 1
	fmt.Println(a10) //1
	//截取某一段，比如34，先左移3位，再右移6位
	var b7 uint8 = 20
	b8 := b7 << 3 >> 6
	fmt.Println(b8, strconv.FormatInt(int64(b8), 2)) //2 10
	//循环右移3位
	a11 := a<<5 + a>>3
	fmt.Println(a11, strconv.FormatInt(int64(a11), 2)) //130 10000010
	//循环左移3位
	a12 := a>>5 + a<<3
	fmt.Println(a12, strconv.FormatInt(int64(a12), 2)) //160 10100000
}

//所谓的补码，比如-5，为10000101，除符号位外所有位取反 11111010
//加00000001 11111011，所以-5的补码是11111011
func b4() {
	var a1 uint16 = 12757
	fmt.Println(strconv.FormatInt(int64(a1), 2)) //0011000111010101
	//获取-12757的补码
	a2 := (^(a1 << 1) >> 1) | (1 << 15) + 1
	fmt.Println(a2, strconv.FormatInt(int64(a2), 2)) //52779 1100111000101011
	//还原-12757补码
	a3 := (^(a2 << 1) >> 1) | (1 << 15) + 1
	fmt.Println(a3, strconv.FormatInt(int64(a3), 2)) //45525 1011000111010101
	//获取整数值
	a4 := a3 &^ (1 << 15)
	fmt.Println(a4, a4 == a1) //12757 true
	//测试一下能不能直接用binary接口来解补码？可以，但是字长必须要贴合值大小，不能过大
	var b1 uint16 = 52779
	b4 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b4, b1)
	fmt.Println(int16(binary.LittleEndian.Uint16(b4))) //-12757
	var b5 uint8 = 250
	fmt.Println(int8(b5), int64(int8(b5))) //-6 -6
	var b6 uint16 = 250
	fmt.Println(int16(b6), int8(b6), int64(int8(b6))) //250 -6 -6，这里用int16解析int8就无效了
	var b7 = byte(251)
	fmt.Println(b7, int8(b7), int64(int8(b7))) //251 -5 -5
	//这个解析接口ParseInt并不能拆补码
	b2, _ := strconv.ParseUint("1100111000101011", 2, 16)
	b3, _ := strconv.ParseInt("1100111000101011", 2, 16)
	fmt.Println(b2, b3) //52779 32767
}

//求一个数的二进制位数
func getBit(n float64) float64 {
	return math.Floor((math.Log(n) / math.Log(2)) + 1)
}

func b5() {
	a, b := 1, 0
	fmt.Println(a&a, a&b, b&a, b&b) //1 0 0 0
	fmt.Println(a|a, a|b, b|a, b|b) //1 1 1 0
	fmt.Println(a^a, a^b, b^a, b^b) //0 1 1 0
	//无符号数全部置1后变负数，有符号数则减1后变正数，2是10，5是101
	fmt.Println(^2, ^-2, ^5, ^-5) //-3 1 -6 4
	fmt.Println(strconv.FormatInt(^2, 2), strconv.FormatInt(^-2, 2),
		strconv.FormatInt(^5, 2), strconv.FormatInt(^-5, 2))
	//-11 1 -110 100
}
