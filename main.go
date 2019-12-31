package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/linnv/logx"

	"github.com/gin-gonic/gin"
	"github.com/linnv/simdog/metric"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"qnmock/api"
	conf "qnmock/config"
	"qnmock/internal/grpc"
)

func Mid() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		metric.IncQps()
	}
}

func main() {
	conf.InitFlag()
	flag.Parse()
	config := conf.InitConfig()

	api.Init()
	grpc.Init()

	// gin.SetMode(gin.ReleaseMode)
	exit := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	router := gin.Default()

	appName := config.AppName
	prometheus.MustRegister(metric.NewCpuMetric(appName))
	prometheus.MustRegister(metric.NewNetMetric(appName))
	prometheus.MustRegister(metric.NewQpsMetric(appName))

	//monitor
	router.Use(Mid())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/metrics", gin.WrapH(promhttp.Handler()))

	router.POST("/gcable", api.Hello)

	port := ":" + config.ServerPort
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Print("closing gin server\n")
			} else {
				panic(err)
			}
		}
	}()

	grpcServer := grpc.StartEngine(exit)

	log.Printf("gin run on port %s\n", port)
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	log.Print("use c-c to exit: \n")
	<-sigChan

	grpcServer.GracefulStop()
	close(exit)
	ctx := context.Background()
	if err := server.Shutdown(ctx); err != nil {
		logx.Errorf("err: %+v\n", err)
	}

	wg.Wait()
	os.Exit(0)
}
