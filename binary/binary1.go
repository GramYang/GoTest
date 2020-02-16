package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	//测试uint16切换字节，两个字节最多存储65535
	test1()
}

func test1() {
	var length = 65535
	i := make([]byte, 2)
	binary.LittleEndian.PutUint16(i, uint16(length))
	fmt.Println(i)
	j := binary.LittleEndian.Uint16(i)
	fmt.Println(j)
}
