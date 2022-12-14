package main

/*

Socket概念: Socket是BSD UNIX的进程通信机制，通常也称作”套接字”，用于描述IP地址和端口，是一个通信链的句柄
	1、Socket是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，Socket其实就是一个门面模式
	2、Socket把复杂的TCP/IP协议族隐藏在Socket后面(可以理解为TCP/IP网络的API)，它定义了许多函数或例程，
		对用户来说只需要调用Socket规定的相关函数，让Socket去组织符合指定的协议数据然后进行通信
		可以用它们来开发TCP/IP网络上的应用程序，应用程序通过”Socket套接字”向网络发出请求或者应答网络请求
		
Socket类型:
	1、流式Socket
		流式是一种面向连接的Socket，针对于面向连接的TCP服务应用
		TCP：比较靠谱，面向连接，比较慢
		TCP就像货到付款的快递，送到家还必须见到你人才算一整套流程
	2、数据报式Socket
		数据报式Socket是一种无连接的Socket，针对于无连接的UDP服务应用
		UDP：不是太靠谱，比较快
		UDP就像某快递快递柜一扔就走管你收到收不到，一般直播用UDP
*/
