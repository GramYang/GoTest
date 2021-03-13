package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	//小端和大端的方法最少要写2个字节
	t1()
}

func t1() {
	var l = 65535
	i := make([]byte, 2)
	binary.LittleEndian.PutUint16(i, uint16(l))
	fmt.Println(i) //[255 255]
	j := binary.LittleEndian.Uint16(i)
	fmt.Println(j) //65535
	var l1 = 255
	i1 := make([]byte, 2)
	i2 := make([]byte, 2)
	//这种方法最少还是写了2个字节，这是专门用于64位的整数
	fmt.Println(binary.PutUvarint(i1, uint64(l1))) //2
	fmt.Println(binary.PutVarint(i2, int64(l1)))   //2
	v1, v2 := binary.Uvarint(i1)
	v3, v4 := binary.Varint(i2)
	fmt.Println(v1, v2) //255 2
	fmt.Println(v3, v4) //255 2
	//单个字节直接用byte强转即可
	b1 := byte(l1)
	fmt.Println(b1, int(b1)) //255 255
	//byte和bool不能互相转换，直接用int的1和0来表示bool吧
	b2 := byte(1)
	b3 := int(b2)
	fmt.Println(b2, b3 == 1) //1 true
	//越界测试，会截取低位
	var l2 uint16 = 12757
	b4 := byte(l2)
	fmt.Println(b4, strconv.FormatInt(12757, 2), strconv.FormatInt(213, 2))
	//213 11000111010101 11010101
	bs1 := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs1, 12757)
	fmt.Println(bs1[1], strconv.FormatInt(12757, 2), strconv.FormatInt(49, 2))
	//49 11000111010101 110001
}
