package junit

/*
Go语言中的测试依赖go test命令, go test命令是一个按照一定约定和组织的测试代码的驱动程序。
在包目录内，所有以 _test.go 为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。

在*_test.go文件中有三种类型的函数:
    类型	    格式	              作用
    测试函数	函数名前缀为Test	    测试程序的一些逻辑行为是否正确
    基准函数	函数名前缀为Benchmark	测试函数的性能
    示例函数	函数名前缀为Example	    为文档提供示例文档

    go test命令会遍历所有的 *_test.go文件中 符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。

测试函数常用命令：
    1、直接运行go test:只会输出测试用例失败的，成功的不输出
    2、go test -v: 不管成功还是失败的都会输出
    3、go test -run="测试用例函数名称" : 过滤出命中的测试函数进行测试
    4、go test -cover: 查看测试代码覆盖率

测试函数注意事项：
    1、文件名必须以xx_test.go命名
    2、方法必须是Test[^a-z]开头,Test后面跟的第一个字母必须大写，不然识别不到
    3、方法参数必须 t *testing.T
    4、使用go test执行单元测试
    
*/

import (
    "reflect"
    "testing"
    "fmt"
    "os"
)

/*
    测试函数示例，基本的功能性测试
    测试函数名必须以Test开头，必须接收一个*testing.T类型参数 
*/

//基础测试示例
func TestSplit(t *testing.T) {
    got := Split("a:b:c", ":")         // 程序输出的结果
    want := []string{"a", "b", "c"}    // 期望的结果
	// 因为slice不能比较直接，借助反射包中的方法比较
    if !reflect.DeepEqual(want, got) { 
        t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
    }
}

/*
测试组：一次执行多个测试用例
问题：是没办法一眼看出来具体是哪个测试用例失败了
*/
func TestGroupSplit(t *testing.T) {
    // 定义一个测试用例类型
    type test struct {
        input string
        sep   string
        want  []string
    }
    // 定义一个存储测试用例的切片
    tests := []test{
        {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
        {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
        {input: "abcd", sep: "bc", want: []string{"a", "d"}},
        {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
    }
    // 遍历切片，逐一执行测试用例
    for _, tc := range tests {
        got := Split(tc.input, tc.sep)
        if !reflect.DeepEqual(got, tc.want) {
            t.Errorf("excepted:%v, got:%v", tc.want, got)
        }
    }
}

/*
子测试：一次执行多个测试用例
    相比于测试组方式，是解决了测试用例失败容易定位的问题，
    在测试组的机构基础上使用map结构，对每个测试用例添加一个key，并打印
*/
func TestOfSubSplit(t *testing.T) {
    type test struct { // 定义test结构体
        input string
        sep   string
        want  []string
    }
    tests := map[string]test{ // 测试用例使用map存储
        "simple": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
        "wrong sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
        "more sep": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
        "错误用例leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
    }
    for name, tc := range tests {
        got := Split(tc.input, tc.sep)
        if !reflect.DeepEqual(got, tc.want) {
            // 将测试用例的name格式化输出,一眼就能看出失败的测试用例
            t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got) 
        }
    }
}
//Go1.7+中新增了子测试，使用t.Run执行子测试，比自己实现的子测试信息更全：
func TestOfGoSubSplit(t *testing.T) {
    type test struct { // 定义test结构体
        input string
        sep   string
        want  []string
    }
    tests := map[string]test{ // 测试用例使用map存储
        "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
        "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
        "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
        "错误用例leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
    }
    for name, tc := range tests {
        //使用t.Run()执行子测试, t.Run()接受一个测试用例名称和一个对应的匿名测试函数
        t.Run(name, func(t *testing.T) { 
            got := Split(tc.input, tc.sep)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("excepted:%#v, got:%#v", tc.want, got)
            }
        })
    }
}


/*
Setup与TearDown
测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）

TestMain
通过在*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。
如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。
TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和拆卸（teardown）。
退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit
*/
//使用TestMain来设置Setup和TearDown的示例如下
func TestMain(m *testing.M) {
    fmt.Println("write setup code here...") // 测试之前的做一些设置
    // 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
    retCode := m.Run()                         // 执行测试
    fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
    os.Exit(retCode)                           // 退出测试
} 


/*
子测试的Setup与Teardown
为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。
*/
//下面我们定义两个函数工具函数如下:
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
    t.Log("如有需要在此执行:测试之前的setup")
    return func(t *testing.T) {
        t.Log("如有需要在此执行:测试之后的teardown")
    }
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
    t.Log("如有需要在此执行:子测试之前的setup")
    return func(t *testing.T) {
        t.Log("如有需要在此执行:子测试之后的teardown")
    }
}
//使用
func TestSubSetupAndTeardownSplit(t *testing.T) {
    type test struct { // 定义test结构体
        input string
        sep   string
        want  []string
    }
    tests := map[string]test{ // 测试用例使用map存储
        "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
        "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
        "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
        "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"", "枯藤", "树昏鸦"}},
    }
    teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
    defer teardownTestCase(t)            // 测试之后执行testdoen操作

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
            teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
            defer teardownSubTest(t)           // 测试之后执行testdoen操作
            got := Split(tc.input, tc.sep)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("excepted:%#v, got:%#v", tc.want, got)
            }
        })
    }
} 