package crawler_client

import (
	"google.golang.org/grpc"
	"taylors_proto/taylors_stock"
)

var Grpc_cli taylors_stock.ServiceClient

func init() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(basics.Conf.CrawlerAddr, grpc.WithInsecure())
	if err != nil {
		logger.Logger.Errorln("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	//defer conn.Close()

	// 创建Waiter服务的客户端
	Grpc_cli = taylors_stock.NewServiceClient(conn)
}
