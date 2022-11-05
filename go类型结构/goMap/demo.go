package main

import "fmt"

/*
map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

语法定义：map[KeyType]ValueType
	KeyType:表示键的类型。
    ValueType:表示键对应的值的类型。
*/

func main() {
	//基本使用
	scoreMap := make(map[string]int, 8)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    fmt.Println(scoreMap)
    fmt.Println(scoreMap["小明"])
    fmt.Printf("type of a:%T\n", scoreMap)

	//判断元素是否存在
	v, ok := scoreMap["张三"]
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("查无此人")
    }

	//删除元素
	delete(scoreMap, "小明")

	//遍历元素
	for k, v := range scoreMap {
        fmt.Println(k, v)
    }

	//声明时，初始化元素
	userInfo := map[string]string{
        "username": "pprof.cn",
        "password": "123456",
    }
    fmt.Println(userInfo)
}