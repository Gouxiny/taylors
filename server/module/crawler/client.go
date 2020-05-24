package crawler

import (
	"google.golang.org/grpc"
	"taylors_proto/taylors_stock"
)

var Grpc_cli taylors_stock.ServiceClient

func init() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// 函数结束时关闭连接
	//defer conn.Close()

	// 创建Waiter服务的客户端
	Grpc_cli = taylors_stock.NewServiceClient(conn)
}
