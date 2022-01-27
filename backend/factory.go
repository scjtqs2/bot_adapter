package backend

import (
	"github.com/karlseguin/ccache/v2"
	"github.com/scjtqs2/bot_adapter/adapter"
	auth2 "github.com/scjtqs2/bot_adapter/auth"
	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/lru"
	"github.com/scjtqs2/bot_adapter/pb/service"
	"go.uber.org/dig"
)

// ServiceFactory 初始化配置
func ServiceFactory(ct *dig.Container) (service.AdapterServiceServer, *dig.Container) {
	var cfg *config.Config
	var svc service.AdapterServiceServer
	err := ct.Invoke(func(opt *config.Config) {
		cfg = opt
	})
	if err != nil {
		panic(err)
	}
	cache := lru.NewCache()
	_ = ct.Provide(func() *ccache.Cache {
		return cache
	})
	// 鉴权服务
	au, _ := auth2.NewAuthServer(ct)
	_ = ct.Provide(func() *auth2.AuthServer {
		return au
	})
	adapt, _ := adapter.NewAdapter(cfg)
	_ = ct.Provide(func() adapter.AdapterInterface {
		return adapt
	})
	bkd := EventResoveBackend{
		EventChan: adapt.GetChan(),
		Config:    cfg,
	}
	_ = ct.Provide(func() *EventResoveBackend {
		return &bkd
	})
	s, err := NewAdapterService(ct) // 最内层
	if err != nil {
		panic(err)
	}
	_ = ct.Provide(func() *AdapterService {
		return s
	})
	s2, err := NewPermissionService(ct, s) // 校验完auth 走到 权限校验
	if err != nil {
		panic(err)
	}
	svc, err = NewAuthorizedService(ct, s2) // 最外层是auth
	if err != nil {
		panic(err)
	}
	return svc, ct
}
