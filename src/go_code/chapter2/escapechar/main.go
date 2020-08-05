package main

import "fmt" //fmt包中主要提供格式化和输出，输入的函数
func main() { //演示转义字符的使用
	fmt.Println("tom\tjack")
	fmt.Println("tomjack") // \t tab

	fmt.Println("hello\nWorld")
	fmt.Println("C:\\Users\\tianyzho\\WorkSpace")
	fmt.Println("tom说\"nihao\"")
	fmt.Println("天龙八部雪山飞狐\r张飞") //当前行顶头开始输出，覆盖掉以前的内容， result ： 张飞八部雪山飞狐

	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}
