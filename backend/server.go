package backend

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/pb/service"
)

type GrpcServer struct {
	config  *config.Config
	service service.AdapterServiceServer
	Event   *EventResoveBackend
}

// NewServer 初始化grpc服务
func NewServer(ct *dig.Container) (*GrpcServer, error) {
	var g GrpcServer
	g.service, ct = ServiceFactory(ct)
	_ = ct.Invoke(func(opt *config.Config) {
		g.config = opt
	})
	_ = ct.Invoke(func(e *EventResoveBackend) {
		g.Event = e
	})
	return &g, nil
}

// Servc 启动 grpc服务
func (g *GrpcServer) Servc() {
	go g.Event.DoEvent()

	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAgeGrace: time.Duration(g.config.GrpcConfig.MaxConnectionAgeGrace) * time.Second,
		MaxConnectionIdle:     time.Duration(g.config.GrpcConfig.MaxConnectionIdle) * time.Second,
		Timeout:               time.Duration(g.config.GrpcConfig.Timeout) * time.Second,
	}))
	service.RegisterAdapterServiceServer(s, g.service)
	go func() {
		listen, err := net.Listen("tcp", g.config.GrpcConfig.Addr)
		if err != nil {
			log.Fatalf("tcp listen failed:%v", err)
		}
		err = s.Serve(listen)
		if err != nil {
			log.Fatalf("开启服务失败: %v", err)
			return
		}
	}()
}
