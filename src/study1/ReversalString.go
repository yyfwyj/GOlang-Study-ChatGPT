package study1

import "fmt"

func ReversalString(str string) string {
	// 判断数据长度
	if len(str) == 0 {
        fmt.Print("123")
	}
	// 处理数据
	runes := []rune(str)
    /*
        
    */
	for i, j := 0, len(runes)-1; i < len(runes)/2; i++ {

		fmt.Println(string(runes[i]), i,"====================> runes[i]")
		fmt.Println(string(runes[j-i]), j-i,"====================> runes[j-i]")

		runes[i], runes[j-i] = runes[j-i], runes[i]
	}
	return string(runes)
}
