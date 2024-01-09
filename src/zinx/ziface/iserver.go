package ziface

type IServer interface {
	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Server 开始业务
	Server()
}
