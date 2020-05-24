package crawler

import (
	"google.golang.org/grpc"
	"taylors/global"
	"taylors_proto/taylors_stock"
)

var Grpc_cli taylors_stock.ServiceClient

func Init() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(global.GVA_CONFIG.Module.Crawler, grpc.WithInsecure())
	if err != nil {
		global.GVA_LOG.Error("grpc 连接失败:", err)
		return
	}
	// 函数结束时关闭连接
	//defer conn.Close()

	// 创建Waiter服务的客户端
	Grpc_cli = taylors_stock.NewServiceClient(conn)
}
