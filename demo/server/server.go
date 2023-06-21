package server

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func server1() {
	http.HandleFunc("/server1/hello", helloHandler)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "server1 Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func epollServer() {
	// 创建监听套接字
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	//// 创建 epoll 实例
	//epollfd, err := syscall.EpollCreate1(0)
	//if err != nil {
	//	panic(err)
	//}
	//defer syscall.Close(epollfd)
	//
	//// 将监听套接字添加到 epoll 中
	//event := syscall.EpollEvent{Events: syscall.EPOLLIN, Fd: int32(l.(*net.TCPListener).Fd())}
	//if err := syscall.EpollCtl(epollfd, syscall.EPOLL_CTL_ADD, int(l.(*net.TCPListener).Fd()), &event); err != nil {
	//	panic(err)
	//}
	//
	//// 处理事件
	//events := make([]syscall.EpollEvent, 10)
	//for {
	//	n, err := syscall.EpollWait(epollfd, events, -1)
	//	if err != nil {
	//		panic(err)
	//	}
	//	for i := 0; i < n; i++ {
	//		if events[i].Fd == int32(l.(*net.TCPListener).Fd()) {
	//			// 处理新连接
	//			conn, err := l.Accept()
	//			if err != nil {
	//				fmt.Println("accept error:", err)
	//				continue
	//			}
	//			fmt.Println("new connection:", conn.RemoteAddr())
	//			// 将连接套接字添加到 epoll 中
	//			event := syscall.EpollEvent{Events: syscall.EPOLLIN | syscall.EPOLLET, Fd: int32(conn.(*net.TCPConn).Fd())}
	//			if err := syscall.EpollCtl(epollfd, syscall.EPOLL_CTL_ADD, int(conn.(*net.TCPConn).Fd()), &event); err != nil {
	//				fmt.Println("epoll_ctl error:", err)
	//				syscall.Close(int(conn.(*net.TCPConn).Fd()))
	//				continue
	//			}
	//		} else {
	//			// 处理已连接套接字
	//			conn := net.Conn{syscall.RawConn{events[i].Fd}}
	//			buf := make([]byte, 1024)
	//			n, err := conn.Read(buf)
	//			if err != nil {
	//				fmt.Println("read error:", err)
	//				continue
	//			}
	//			fmt.Println("read from", conn.RemoteAddr(), ":", string(buf[:n]))
	//			// 将数据写回客户端
	//			_, err = conn.Write([]byte("Hello, World!"))
	//			if err != nil {
	//				fmt.Println("write error:", err)
	//				continue
	//			}
	//		}
	//	}
	//}
}

func selectServer() {

	// 创建监听套接字
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// 创建通道
	connections := make(chan net.Conn)

	// 启动协程处理连接请求
	go func() {
		for {
			// 接受连接请求
			conn, err := l.Accept()
			if err != nil {
				panic(err)
			}
			// 将连接请求发送到通道中
			connections <- conn
		}
	}()

	// 处理连接请求
	for {
		select {
		case conn := <-connections:
			// 处理连接请求
			go selectHandleConnection(conn)
		}
	}
}

// 处理连接请求
func selectHandleConnection(conn net.Conn) {
	defer conn.Close()
	// 处理请求
	conn.Write([]byte("select server,hello world"))
}
