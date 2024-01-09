package znet

import (
	"fmt"
	"net"
	"zinx-demo/src/zinx/ziface"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 当前连接的ID 也可以称作SessionID， ID全局唯一
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool
	// 当前连接的处理方法
	handleAPI ziface.HandFunc
	// 告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callbackApi ziface.HandFunc) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    callbackApi,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("")
}
