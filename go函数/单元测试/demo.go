package junit

import "strings"

func Split(s, sep string) (result []string) {
    i := strings.Index(s, sep)

    for i > -1 {
        result = append(result, s[:i])
        s = s[i+len(sep):]
        i = strings.Index(s, sep)
    }
    result = append(result, s)
    return
}

// 优化后的函数，实现通过make分配好容量足够大内存
func SplitOptimalize(s, sep string) (result []string) {
    result = make([]string, 0, strings.Count(s, sep)+1)
    i := strings.Index(s, sep)
    for i > -1 {
        result = append(result, s[:i])
        s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
        i = strings.Index(s, sep)
    }
    result = append(result, s)
    return
}


// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
    if n < 2 {
        return n
    }
    return Fib(n-1) + Fib(n-2)
} 