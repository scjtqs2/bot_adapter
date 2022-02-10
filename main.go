package main

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/scjtqs2/bot_adapter/backend"
	"github.com/scjtqs2/bot_adapter/config"
)

// Version 版本信息
var Version = "devel"

func main() {
	log.Printf("欢迎使用 bot_adapter for cqhttp. ver=%s", Version)
	ct := bootStrap()
	svc, _ := backend.NewServer(ct)
	svc.Servc()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	os.Exit(1)
}

func bootStrap() *dig.Container {
	cfg := config.Parse()
	_ = cfg.Save()
	ct := dig.New()
	_ = ct.Provide(func() *config.Config {
		return cfg
	})
	return ct
}
