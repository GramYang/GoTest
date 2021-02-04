package main

import (
	"fmt"
	"regexp"
)

//正则表达式语法，参考：
//https://zhuanlan.zhihu.com/p/28672572
//https://www.cnblogs.com/sunsky303/p/11051468.html

//转义字符列表
//\n 换行 \r 回车 \t tab \v 垂直tab \w 字母、数字、下划线 \W \w的反例 \s 空格、tab \S \s反例
//\d 数组 \D \d反例 \b单词的边界 \B \b反例 \\ 斜杠\

//集合
//[abc] a或b或c
//[^abc] 匹配abc之外的任意字符

//量词
//a{2} 匹配n次，比如a{2}，匹配aa
//{m, n} 匹配m-n次，优先匹配n次，比如a{1,3}，可以匹配aaa、aa、a
//{m,} 匹配m-∞次，优先匹配∞次，比如a{1,}，可以匹配aaaa...
//? 匹配0次或1次，优先匹配1次，相当于{0,1}
//+ 匹配1-n次，优先匹配n次，相当于{1,}
//* 匹配0-n次，优先匹配n次，相当于{0,}
//a{1, 3} 匹配字符串'aaa'的话，会匹配aaa而不是a
//a{1, 3}? 匹配字符串'aaa'的话，会匹配a而不是aaa

//字符边界
//^在[]外表示匹配开头的意思
//^abc 可以匹配abc，但是不能匹配aabc
//$表示匹配结尾的意思
//abc$ 可以匹配abc，但是不能匹配abcc
//上面提到的\b表示单词的边界
//abc\b 可以匹配 abc ，但是不能匹配 abcc

//选择表达式
//正则中用|来表示分组，a|b表示匹配a或者b的意思
//123|456|789 // 匹配 123 或 456 或 789

//分组与引用
//分组的语法是圆括号包裹(xxx)
//(abc){2} 匹配abcabc
//分组不能放在[]中，分组中还可以使用选择表达式
//(123|456){2} 匹配 123123、456456、123456、456123
//在分组的(后面添加?:可以让分组变为非捕获分组，非捕获分组可以提高性能和简化逻辑
//'123'.match(/(?123)/) 返回 ['123']
//'123'.match(/(123)/)  返回 ['123', '123']
//引用的语法是\数字，数字代表引用前面第几个捕获分组，注意非捕获分组不能被引用
//<([a-z]+)><\/\1> 可以匹配 `<span></span>` 或 `<div></div>`等

//预搜索
//(?=1)2 可以匹配12，不能匹配22
//(?!1)2 可有匹配22，不能匹配12

//修饰符
///xxx/gi 最后面的g和i就是两个修饰符
//g正则遇到第一个匹配的字符就会结束，加上全局修复符，可以让其匹配到结束
//i正则默认是区分大小写的，i可以忽略大小写
//m正则默认遇到换行符就结束了，不能匹配多行文本，m可以让其匹配多行文本

func main() {
	//regexp正则表达式解析包使用测试
	rt1()
}

func rt1() {
	text := `Hello 世界！123 Go.`
	reg := regexp.MustCompile(`[a-z]+`)               // 找出所有小写字母的子串
	fmt.Printf("1 %q\n", reg.FindAllString(text, -1)) //["ello" "o"]
	reg = regexp.MustCompile(`[^a-z]+`)               //找出所有非小写字母的子串
	fmt.Printf("2 %q\n", reg.FindAllString(text, -1)) //["H" " 世界！123 G" "."]
	reg = regexp.MustCompile(`[\w]+`)
	fmt.Printf("3 %q\n", reg.FindAllString(text, -1)) //["Hello" "123" "Go"]
	reg = regexp.MustCompile(`[^\w\s]+`)
	fmt.Printf("4 %q\n", reg.FindAllString(text, -1)) //["世界！" "."]
	reg = regexp.MustCompile(`[[:upper:]]+`)
	fmt.Printf("5 %q\n", reg.FindAllString(text, -1)) //["H" "G"]
	reg = regexp.MustCompile(`[[:^ascii:]]+`)
	fmt.Printf("6 %q\n", reg.FindAllString(text, -1)) //["世界！"]
	reg = regexp.MustCompile(`[\pP]+`)
	fmt.Printf("7 %q\n", reg.FindAllString(text, -1)) //["！" "."]
	reg = regexp.MustCompile(`[\PP]+`)
	fmt.Printf("8 %q\n", reg.FindAllString(text, -1)) //["Hello 世界" "123 Go"]
	reg = regexp.MustCompile(`[\p{Han}]+`)
	fmt.Printf("9 %q\n", reg.FindAllString(text, -1)) //["世界"]
	reg = regexp.MustCompile(`[\P{Han}]+`)
	fmt.Printf("10 %q\n", reg.FindAllString(text, -1)) //["Hello " "！123 Go."]
	reg = regexp.MustCompile(`Hello|Go`)
	fmt.Printf("11 %q\n", reg.FindAllString(text, -1)) //["Hello" "Go"]
	reg = regexp.MustCompile(`^H.*\s`)
	fmt.Printf("12 %q\n", reg.FindAllString(text, -1)) //["Hello 世界！123 "]
	reg = regexp.MustCompile(`(?U)^H.*\s`)
	fmt.Printf("13 %q\n", reg.FindAllString(text, -1)) //["Hello "]
	reg = regexp.MustCompile(`(?i:^hello).*Go`)
	fmt.Printf("14 %q\n", reg.FindAllString(text, -1)) //["Hello 世界！123 Go"]
	reg = regexp.MustCompile(`\QGo.\E`)
	fmt.Printf("15 %q\n", reg.FindAllString(text, -1)) //["Go."]
	reg = regexp.MustCompile(`(?U)^.* `)
	fmt.Printf("16 %q\n", reg.FindAllString(text, -1)) //["Hello "]
	reg = regexp.MustCompile(` [^ ]*$`)
	fmt.Printf("17 %q\n", reg.FindAllString(text, -1)) //[" Go."]
	reg = regexp.MustCompile(`(?U)\b.+\b`)
	fmt.Printf("18 %q\n", reg.FindAllString(text, -1)) //["Hello" " 世界！" "123" " " "Go"]
	reg = regexp.MustCompile(`[^ ]{1,4}o`)
	fmt.Printf("19 %q\n", reg.FindAllString(text, -1)) //["Hello" "Go"]
	reg = regexp.MustCompile(`(?:Hell|G)o`)
	fmt.Printf("20 %q\n", reg.FindAllString(text, -1)) //["Hello" "Go"]
	reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
	fmt.Printf("21 %q\n", reg.ReplaceAllString(text, "$3$2$1")) //"Go 世界！123 Hello."
}
