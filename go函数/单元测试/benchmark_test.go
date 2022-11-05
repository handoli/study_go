package junit

/*
基准函数注意事项：
	1、文件名必须以xx_test.go命名
	2、方法必须是Benchmark[^a-z]开头,Benchmark后面跟的第一个字母必须大写，不然识别不到
	3、方法参数必须 b *testing.B
	4、通过执行go test -bench="包含的测试基准函数名称"命令执行基准测试
*/

import (
	// "reflect"
	"testing"
	"time"
)

//基准函数测试示例
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}
/*
执行结果:
	handoli@handoli 单元测试 % go test -bench=Split -benchmem
	goos: darwin
	goarch: amd64
	cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
	BenchmarkSplit-8  10135219   116.7 ns/op   48 B/op   2 allocs/op
	PASS
	ok      _/Users/handoli/_summary/example/gomod/study_go/go函数/单元测试 1.699s

命令解释:
	-benchmem参数，来获得内存分配的统计数据
结果解释:
	BenchmarkSplit-8: 表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。
	10135219和116.7 ns/op: 表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。
	48 B/op: 表示每次操作内存分配了48字节
	2 allocs/op: 则表示每次操作进行了2次内存分配
*/


//优化后的 基准函数测试
func BenchmarkSplitOptimalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitOptimalize("枯藤老树昏鸦", "老")
	}
}
/*
执行结果:
	goos: darwin
	goarch: amd64
	cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
	BenchmarkSplitOptimalize-8      12074641                93.94 ns/op           32 B/op          1 allocs/op
	PASS
	ok      _/Users/handoli/_summary/example/gomod/study_go/go函数/单元测试 1.491s
优化结果解释:
	优化后的内存只分配了一次，比没优化时少一次
	优化后的内存分配大小为32，没优化时的分配为48

*/


/*
性能比较函数
	基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，
	比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？
	再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试

性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用：
	func benchmark(b *testing.B, size int){ ... }
	func Benchmark10(b *testing.B){ benchmark(b, 10) }
	func Benchmark100(b *testing.B){ benchmark(b, 100) }
	func Benchmark1000(b *testing.B){ benchmark(b, 1000) } 
*/
func benchmarkFib(b *testing.B, n int) {
    for i := 0; i < b.N; i++ {
        Fib(n)
    }
}
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) } 
/*
执行结果：
	handoli@handoli 单元测试 % go test -bench=Fib
	goos: darwin
	goarch: amd64
	cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
	BenchmarkFib1-8         860522644                1.381 ns/op
	BenchmarkFib2-8         287761942                4.208 ns/op
	BenchmarkFib3-8         173010430                7.046 ns/op
	BenchmarkFib10-8         4606314               249.4 ns/op
	BenchmarkFib20-8           38403             31065 ns/op
	BenchmarkFib40-8               3         457563857 ns/op
	PASS
	ok      _/Users/handoli/_summary/example/gomod/study_go/go函数/单元测试 10.688s

	默认情况下，每个基准测试至少运行1秒。
	如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行

最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。
像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果
如下：
	handoli@handoli 单元测试 % go test -bench=Fib40 -benchtime=20s
	goos: darwin
	goarch: amd64
	cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
	BenchmarkFib40-8              51         471979580 ns/op
	PASS
	ok      _/Users/handoli/_summary/example/gomod/study_go/go函数/单元测试 24.805s
*/


/*
重置时间
b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作
*/
func BenchmarkResetTimeSplit(b *testing.B) {
    time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
    b.ResetTimer()              // 重置计时器
    for i := 0; i < b.N; i++ {
        Split("枯藤老树昏鸦", "老")
    }
}

/*
并行测试	以并行的方式执行给定的基准测试
RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行，其中goroutine数量的默认值为GOMAXPROCS。
用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。
RunParallel通常会与-cpu标志一同使用
*/
func BenchmarkSplitParallel(b *testing.B) {
    // b.SetParallelism(1) // 设置使用的CPU数
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            Split("枯藤老树昏鸦", "老")
        }
    })
}
/*
执行结果：
	handoli@handoli 单元测试 % go test -bench=SplitParallel
	goos: darwin
	goarch: amd64
	cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
	BenchmarkSplitParallel-8        26789872                41.20 ns/op
	PASS
	ok      _/Users/handoli/_summary/example/gomod/study_go/go函数/单元测试 2.176s
*/