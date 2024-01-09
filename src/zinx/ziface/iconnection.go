package ziface

import "net"

type IConnection interface {

	// Start 启动连接，让当前连接开始工作
	Start()

	// Stop 停止连接，结束当前连接状态
	Stop()

	// 从当前连接获取原始的socket TCPConn GetTCPConnection() *net.TCPConn 连接id
	GetConnID() uint32
}

// 定义一个同一处理链接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
