package junit

import(
	"fmt"
)

/*
示例函数:
	1、文件名必须以xx_test.go命名
	2、方法必须是Example[^a-z]开头,Example后面跟的第一个字母必须大写，不然识别不到
	3、既没有参数也没有返回值

编写代码示例代码三个用处：
	1、示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
    2、示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。
        split $ go test -run Example
        PASS
        ok      github.com/pprof/studygo/code_demo/test_demo/split       0.006s
    3、示例函数提供了可以直接运行的示例代码，
	   可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。
*/

func ExampleSplit() {
    fmt.Println(Split("a:b:c", ":"))
    fmt.Println(Split("枯藤老树昏鸦", "老"))
    // Output:
    // [a b c]
    // [ 枯藤 树昏鸦]
}