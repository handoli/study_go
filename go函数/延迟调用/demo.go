package main

import (
	"fmt"
	"os"
)

/*
defer 后面会接受一个函数，但该函数不会立刻被执行，而是等到包含它的程序返回时(包含它的函数执行了return语句、运行到函数结尾自动返回、对应的goroutine panic），defer函数才会被执行。
	
通常用于资源释放、打印日志、异常捕获等
*/

//defer的基本使用，defer声明的位置（关键）
func OpenFile()  {
	file, err := os.Open("info.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 这里将 defer 放在 err 判断的后面，而不是 os.Open() 的后面
	// 因为若 err != nil ，文件打开是失败的, 没必要释放
	// 若 err != nil, file 有可能为 nil ,这时候释放资源可能会导致程序崩溃
	defer file.Close()
}

/*
defer 的执行时机,defer 后面的函数对外部参数有两种引用方式：
	参数传递：在defer声明时，即将值传递给defer,并缓存起来，调用defer的时候使用缓存值进行计算
	直接引用：根据上下文确定当前值，作用域与外部变量相同
*/
func deferTest() {
	fish := 0	//声明 fish=0
	// 直接引用(闭包),最终会引用外部的fish，打印外部fish
	defer func() {
		fmt.Println("d1: ", fish)
	}()
	// 参数传递,传递时fish=0,打印内部 fish
	defer fmt.Println("d2: ", fish)
	// 参数传递(闭包)，传递fish时 fish=0，fish在内部操作，打印的是内部fish
	defer func(fish int) {
		fish += 2							// fish 只作用于内部
		fmt.Println("d3: ", fish)			//打印内部 fish = 2
	}(fish)							    // 声明时传递, fish = 0
  
	// 直接引用(闭包)，操作的是外部fish
	defer func() {
		fmt.Println("d4: ", fish)			// 打印外部fish
		fish +=2			// fish = 3, 作用域与外部的相同
	}()
  
    //此时fish=1
	fish++
}  


/*
defer 与 return 返回值的关系
有返回值的且带有 defer 函数的方法中， return 语句执行顺序：
	1. 返回值赋值
	2. 调用 defer 函数 (在这里是可以修改返回值的)
	3. return 返回值
*/
// 返回 11
func defer1() (res int) {
	defer func() {
		res ++	// 第二部.res = 10 + 1 = 11
	}()
	return 10	// 第一步.res = 10,  最后一步（第三部）.return res
}

// 返回10
func defer2() (res int) {
	sb := 10	// 第一步. sb = 10
	defer func() {
		sb += 5 // 第三步. sb = 15, 但是 res = 10
	}()
	return sb	// 第二步. res = sb = 10,	最后一步（第四步）. return res(10)
}
//变种案例（未被命名的返回值参数，相当于新参数）
func defer2New() int {
	var i int
	defer func () {
		i++					//引用了外部变量i，修改会影响外部i
		fmt.Println(i)
	}()
	return i  // 注意：未知变量 = i = 0, 因此defer函数中对i的变化不会影响返回值
}

// 返回10
func defer3() (res int) {
	  res = 2	// 第一步. res=2
	  defer func(res int) { 
		  res += 2	// 第三步. 内部res为形参，不影响外边的值 res=2+2=4
		  fmt.Println("内部 res ", res)   // 打印的是内部res=4
	  }(res)		// 值传递 
	  return 10	// 第二步. res = 10,	最后一步（第四步）. return res(10)
  }
  

func main() {
	OpenFile()

	deferTest()

	fmt.Println(defer1())
	
	fmt.Println(defer2())
	fmt.Println(defer2New())

	fmt.Println(defer3())
	
}