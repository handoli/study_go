package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

/*
defer+panic+recover的组合可以发挥出 java 中 try...catch...fanilly 的作用

defer:查看延迟调用示例

panic:内置函数
	1、引发panic的情况有两种:
		一种是程序主动调用panic函数
		另一种是程序产生运行时错误，由运行时检测并抛出
	2、发生panic之后，会终止其后要执行的代码，逐层向上执行函数的defer语句，
	当返回到外层函数时也不会在执行后面的代码，而是逐层向上执行外层函数的defer语句，
	直到运行到函数退出或被recover捕获
	3、panic不但可以在函数正常的流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。defer里面的panic能够被后续执行的defer捕获(必须是同一个gorutine)
	4、可以有连续多个panic被抛出，这种场景只能出现在延迟调用里面，否则不会出现多个panic被抛出的情况。但是只有最后一个panic会被捕获

recover:内置函数
	1、recover()用来捕获panic，阻止panic继续向上传递。recover()和defer一起使用
	2、recover()只有在defer后面的函数体内被直接调用才能有效捕获panic。
		否则返回nil，无法防止panic扩散，异常继续向外传递

注意：panic、recover 参数类型为 interface{}，因此可抛出任何类型对象。
    func panic(v interface{})
    func recover() interface{}
*/

//基本使用
func test() {
    defer func() {
        if err := recover(); err != nil { //recover
            fmt.Println(err.(string)) // 将 interface{} 转型为具体类型。
        }
    }()
    panic("panic error!")  //panic
}  

//向已关闭的通道发送数据会引发panic
func sendDataToChan(){
	// defer中调用recover捕获panic
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan int, 10) //声明chan
	close(ch)	//关闭chan
	ch <- 1		//向关闭的chan发送数据，会引发panic -> send on closed channel
}

//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func morePanic(){
	/*
	捕获最近的一个panic，这里捕获的是defer中的panic("defer--morePanic")
	还往前的panic需要在之前的代码逻辑中加入recover()进行捕获
	如果之前不加recover()进行捕获，则最后的recover()只捕获最后一个panic
	*/
	defer func ()  {
		err := recover()
		fmt.Println(err)
	}()

	defer func(){ //defer中抛出panic
		panic("defer--morePanic")
	}()

	// 捕获最近的一个panic，这里捕获的是panic("morePanic")
	defer func ()  {
		err := recover()
		fmt.Println(err)
	}()

	panic("morePanic")	//手动抛出panic
}

/*
捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。
任何未捕获的错误都会沿调用堆栈向外传递
*/
func effectiveInRecover(){
	defer catch()	//有效！调用的外部函数中的recover可以捕获到调用者函数的panic
	defer func() {
        fmt.Println(recover()) //有效！在defer函数体内直接调用有效
		panic("catch panic")	//这里抛出是为了演示外部函数中的recover捕获panic
    }()
    defer recover()              //无效！直接调用无效，必须是defer函数体内直接调用
    defer fmt.Println(recover()) //无效！当作defer函数的参数传递调用无效
    defer func() {
        func() {
            println("defer inner")
            recover() //无效！defer函数内部的函数内调用无效
        }()
    }()
    panic("test panic")
}

//其他函数的defer中调用外部函数，外部函数中的recover可以捕获到panic异常
func catch(){
	fmt.Println("catch-------",recover())
}

//将代码块重构成匿名函数，如此可确保后续代码被执
//如同main函数中调用的多个函数，每个函数中的panic-recover不影响其他后面函数执行
func codeBlock(x,y int){
	var z int
	fmt.Println("z defult value : ", z)
	func () {
		defer func () {
			if err := recover(); err != nil {
				z = 10
			}
		}()
		panic("codeBlock---panic")
		z = x / y
	}()

	fmt.Printf("x / y = %d\n", z)
}

/*
标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。
通过判断错误对象实例来确定具体错误类型

如何区别使用 panic 和 error 两种方式?
	惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error
*/
var ErrDivByZero = errors.New("division by zero")
func errDemo(x, y int) (int, error) {
    if y == 0 {
        return 0, ErrDivByZero
    }
    return x / y, nil
}

func errNewDemo(){
	defer func() {
        fmt.Println(recover())	//捕获
    }()
	//调用errDemo函数，返回值可能是一个errors类型
    switch z, err := errDemo(10, 0); err {
    case nil:
        println(z)
    case ErrDivByZero:
        panic(err)	//errors类型的通过panic抛出
    }
}

/*
自定义error
*/
//定义error信息结构体
type PathError struct {
    path       string
    op         string
    createTime string
    message    string
}
/*
在调用自定义的Open方法时，需要返回一个error类型的实例
只有我们自定义的PathError实现error接口的Error方法时，PathError才属于error类型
此时返回PathError就满足返回error类型的要求，也就是满足了我们自定义error
*/
func (p *PathError) Error() string {
    return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
        p.op, p.createTime, p.message)
}
/*
Open方法返回error类型
这里可以直接返回os.Open(filename)方法的err
但是我们自定义的PathError实现了error类型，所以也可以返回我们自定义的PathError(也是error类型)
*/
func Open(filename string) error {
    file, err := os.Open(filename) //调用文件系统函数，如果报错，则返回自定义error信息
    if err != nil {
		fmt.Println("当报错时，file对象的状态==",file)
		fmt.Println("error info====",err)
        return &PathError{ //返回自定义的error信息
            path:       filename,
            op:         "read",
			//这里的err.Error()是os.Open(filename)返回的err
            message:    err.Error(),
            createTime: fmt.Sprintf("%v", time.Now()),
        }
    }
    defer file.Close() //关闭资源
    return nil
}


/*
实现类似 try catch 的异常处理
函数参数：
	第一个函数参数：类似于try代码块要处理的逻辑
	第二个函数参数：类似于catch代码块要处理的逻辑，
		注意catch处理的是异常类型，go的异常类型是interface{}
*/
func Try(fun func(), handler func(interface{})) {
	//捕获到panic，交由类似于catch代码块的函数处理
    defer func() {
        if err := recover(); err != nil {
            handler(err)
        }
    }()
	//正常处理逻辑的函数，类似于try块代码的函数，可能会panic
    fun()
}


func main() {
	test()

	sendDataToChan()

	morePanic()

	effectiveInRecover()

	fmt.Println("==================")
	codeBlock(2,1)

	//errors.New
	fmt.Println("===================")
	errNewDemo()

	// try-catch
	fmt.Println("===================")
	Try(func() {	//try逻辑
        panic("try-catch panic")
    }, func(err interface{}) {	//catch逻辑
        fmt.Println(err)
    })

	//自定义Error
	fmt.Println("===================")
	pe := Open("./error.txt")
    switch pe.(type) {
    case *PathError:
		fmt.Println(pe.Error()) //调用我们自定义的PathError实现的error接口的Error方法
    default:
		fmt.Println("无错误，nil")
    }
}
