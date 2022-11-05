package main

import (
	"fmt"
	"net/http"
)

/*
web工作流程:
	1、客户机通过TCP/IP协议建立到服务器的TCP连接
	2、客户端向服务器发送HTTP协议请求包，请求服务器里的资源文档
	3、服务器向客户机发送HTTP协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎
    4、负责处理“动态内容”，并将处理得到的数据返回给客户端
	5、客户机与服务器断开。由客户端解释HTML文档，在客户端屏幕上渲染图形结果

HTTP协议:
	1、超文本传输协议(HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络协议，
	   它详细规定了浏览器和万维网服务器之间互相通信的规则，通过因特网传送万维网文档的数据传送协议
	2、HTTP协议通常承载于TCP协议之上
*/

func main() {
	//http://127.0.0.1:8000/go
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("www.5lmh.com"))
}
