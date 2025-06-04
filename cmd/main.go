package main

import (
	"net/http"
	"time"

	"github.com/lyy42995004/IM-Go/internal/config"
	"github.com/lyy42995004/IM-Go/internal/kafka"
	"github.com/lyy42995004/IM-Go/internal/router"
	"github.com/lyy42995004/IM-Go/internal/server"
	"github.com/lyy42995004/IM-Go/pkg/common/constant"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

func main() {
	defer log.Sync() // 记录日志

	log.Info("start chat server...")

	if  mct := config.GetConfig().MsgChannelType; mct.ChannelType == constant.KAFKA {
		kafka.InitProducer(mct.KafkaTopic, mct.KafkaHosts)
		kafka.InitConsumer(mct.KafkaHosts)
		go kafka.ConsumerMsg(server.ConsumerKafkaMsg)
	}

	newRouter := router.NewRouter()

	go server.MyServer.Start()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        newRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Error("server start error", log.Err(err))
	}
}
