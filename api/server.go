package api

import (

	"github.com/gin-gonic/gin"
	"github.com/chen/simplebank/internal/db"
)
// 服务器处理 银行服务的所有http请求
// *db.Store 是处理客户端的api请求和数据库进行交互
// gin.Engine路由器 帮我们把每个不同的api发送到指定位置（正确的处理程序处理）
type Server struct{
	store *db.Store
	router *gin.Engine
}
// 新的服务器实例，为我们的服务设置所有的http api 路由
func NewServer(store *db.Store) *Server{
	server := &Server{store: store}   //带输入存储的新的服务器
	router := gin.Default()   // 新的路由器
	
	//往路由器中添加路由
	router.POST("/accounts", server.createAccount)    //中间都是中间件，最后一个才是真正的函数
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// 启动http, 在输入地址上运行HTTP服务 开始监听API请求
func(server *Server) Start(address string) error{
	return server.router.Run(address)
}

// gin.H 是索引表 map[string]interface{}
func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}