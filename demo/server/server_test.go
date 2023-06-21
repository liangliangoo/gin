package server

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer1(t *testing.T) {
	server1()
}

func TestEpollServer(t *testing.T) {
	epollServer()
}

func TestSelectServer(t *testing.T) {
	go func() {
		time.Sleep(5 * time.Second)
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}
		// 读取数据
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		// 输出数据
		fmt.Println(string(buf[:n]))
	}()
	selectServer()
}
