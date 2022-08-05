package grpcserv

import (
	"weixin/common/handlers/conf"
	"weixin/common/handlers/log"
	"weixin/common/server/httpserv"
	"weixin/controller"
	"weixin/idl/proto"
	"weixin/libs/officialaccount"
	grpcMiddleware "weixin/middleware/grpc"

	httpMiddleware "weixin/middleware/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"team.wphr.vip/technology-group/infrastructure/rpcserver"
)

var (
	svr = rpcserver.New()

	grpcRegister = func(s *grpc.Server) {
		proto.RegisterWeixinServiceServer(s, new(controller.WeixinService))
	}

	httpRegister = proto.RegisterWeixinServiceHandlerFromEndpoint
)

// Run 启动服务
func Run() error {
	/* 读取conf配置 */
	svr.SetGRPCAddr(conf.Viper.GetString("rpc.grpc_addr"))
	svr.SetHTTPAddr(conf.Viper.GetString("rpc.http_addr"))
	svr.SetGRPCHandlerTimeout(conf.Viper.GetInt("rpc.grpc_handler_timeout"))
	svr.SetHTTPReadTimeout(conf.Viper.GetInt("rpc.http_read_timeout"))
	svr.SetHTTPIdleTimeout(conf.Viper.GetInt("rpc.http_idle_timeout"))
	svr.SetGrpcIdleTimeout(conf.Viper.GetInt("rpc.grpc_idle_timeout"))

	/* 服务注册 */
	svr.SetGRPCRegister(grpcRegister)
	svr.SetHTTPRegister(httpRegister)

	// 【http】recovery中间件
	svr.AddHTTPMiddleware(httpMiddleware.RecoveryWithConfig(httpMiddleware.RecoveryConfig{Log: log.Trace}))
	// 【http】trace中间件
	svr.AddHTTPMiddleware(httpMiddleware.TraceWithConfig(httpMiddleware.TraceConfig{Log: log.Trace}))
	// 【grpc】recovery中间件
	svr.AddGRPCMiddleware(grpcMiddleware.RecoveryInterceptor())
	// 【grpc】trace中间件
	svr.AddGRPCMiddleware(grpcMiddleware.TraceInterceptor())
	// 【grpc】validator中间件, 有校验需求再打开
	svr.AddGRPCMiddleware(grpcMiddleware.ValidatorInterceptor())
	// 【http】auth中间件
	//svr.AddHTTPMiddleware(httpMiddleware.Auth())
	// 【http】page中间件
	svr.AddHTTPMiddleware(httpMiddleware.InitPage())
	// 【http】文件下载中间件
	svr.AddHTTPMiddleware(httpMiddleware.FileDownload())

	/* http 选项 */
	// 过滤http请求的header
	svr.AddHTTPOption(runtime.WithIncomingHeaderMatcher(httpserv.IncomingHeaderMatcher))
	// 过滤http返回的header
	svr.AddHTTPOption(runtime.WithOutgoingHeaderMatcher(httpserv.OutgoingHeaderMatcher))
	// 设置metadata，http req -> grpc metadata
	//svr.AddHTTPOption(runtime.WithMetadata(httpserv.AnnotatorHTTPReq))
	// 设置默认编解码方式
	svr.SetDefaultMarshaler()

	/* grpc-gateway 错误处理 */
	runtime.HTTPError = httpserv.HTTPError

	/* 添加http handle，如上传文件，页面等等 */
	// svr.AddHTTPHandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	//     w.Write([]byte("test"))
	// })

	// 处理微信消息通知
	svr.AddHTTPHandleFunc("/event", officialaccount.ServeWechat)

	return svr.Run()
}
