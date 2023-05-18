/*
	chatGPT每日一题：
	1. chatGPT命令：
	根据我提供的开发语言或者框架的开发规范、开发文档、API、语法以及最流行的知识点，来提供编程开发面试题：
	1.1. 开发语言： Golang
	1.2. 编程开发面试题要求：
	1.2.1 每一种面试题如果提供了最低难度，则提出的编程开发类型面试题不可低于最低难度，如果没有提供，则题目难度随机
	1.2.2 每种语言或者框架出一道面试题，面试题必须是编程开发类型面试题
	1.2.3 根据市面上流行的面试题难度来对你提出的的问题进行分级，分成十级，同时每道题的难度随机，告知我你提出的面试题是几级的
	1.2.4 每道题请尽量是面试可能遇到的面试题以及实际开发过程中容易遇到的编程问题进行提问
	1.2.5 我的Golang水平比较低，golang的题难度请先不要那么大

	题目：
	请编写一个 Golang 程序，实现字符串反转函数，即输入一个字符串，输出其反向字符串。
	
	难度：
	难度级别：4/10
*/

package study1

import "fmt"

func ReversalString(str string) string {
	// 判断数据长度
	if len(str) == 0 {
		return "字符串长度不可为0"
	}

	/*
		处理数据，创建一个 rune切片 接收传入的字符串
		rune 的 作用：
	*/
	runes := []rune(str)

	/*
	        1. 定义一个for循环，遍历传入的字符串切片
			2. 定义 i = 0 j = len(runes) - 1，目的是通过 i 进行 runes 正序的数据获取，通过 j 获得 runes的 反向数据获取
			3. runes[i], runes[j-i] = runes[j-i], runes[i]： 进行同步交换，二者交换数据，进行反向
			4. 注意：
				4.1 i < len(runes)/2 => 该代码是因为 i 只需要遍历前半部分，让 j 遍历后半部分，二者再进行数据交换
				4.2 runes[i], runes[j-i] = runes[j-i], runes[i]： 该代码不像 Javascript，JavaScript的运行逻辑是先赋值前面的，再赋值后面的，而Golang则是同步进行
	*/
	for i, j := 0, len(runes)-1; i < len(runes)/2; i++ {

		fmt.Println(string(runes[i]), i, "====================> runes[i]")
		fmt.Println(string(runes[j-i]), j-i, "====================> runes[j-i]")

		runes[i], runes[j-i] = runes[j-i], runes[i]
	}

	// 返回处理好的反转字符串
	return string(runes)
}
