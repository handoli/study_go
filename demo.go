package main

import (
	"fmt"
)

func assert(i interface{}) {
    value, ok := i.(int)
    fmt.Println(value, ok)
}

func main() {
    var x interface{} = 3
    assert(x)
    var y interface{} = "从0到Go语言微服务架构师"
    assert(y)

    println("123")
}