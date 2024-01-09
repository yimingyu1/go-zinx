package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client Test ... start")
	time.Sleep(3 * time.Second)
	con, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	for {
		_, err := con.Write([]byte("hello ZINX"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := con.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf(" server call back: %s, cnt = %d\n", buf[:cnt], cnt)
		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("[zinx V0.1]")
	go ClientTest()
	s.Server()
}
