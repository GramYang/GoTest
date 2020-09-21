package main

import (
	"context"
	"flag"
	"github.com/devopsfaith/krakend-opencensus"
	"github.com/devopsfaith/krakend-opencensus/exporter"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/influxdb"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/jaeger"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/prometheus"
	_ "github.com/devopsfaith/krakend-opencensus/exporter/zipkin"
	opencensusgin "github.com/devopsfaith/krakend-opencensus/router/gin"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/devopsfaith/krakend/transport/http/client"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "/etc/krakend/configuration.json", "Path to the configuration filename")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	//接收的信号为interrupt和terminated
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//在子goroutine中监视ctx的执行情况
	go func() {
		select {
		case sig := <-sigs: //阻塞等待sigs，如果有则调用cancel中途退出
			log.Println("Signal intercepted:", sig)
			cancel()
		case <-ctx.Done(): //ctx所属的goroutine执行完毕则解除阻塞，这有什么用？
		}
	}()

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")

	// Register stats and trace exporters to export the collected data.
	{
		exporter.Register(logger)

		if err := opencensus.Register(ctx, serviceConfig); err != nil {
			log.Fatal(err)
		}
	}

	bf := func(cfg *config.Backend) proxy.Proxy {
		return proxy.NewHTTPProxyWithHTTPExecutor(cfg, opencensus.HTTPRequestExecutor(client.NewHTTPClient), cfg.Decoder)
	}

	// setup the krakend router
	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         gin.Default(),
		ProxyFactory:   opencensus.ProxyFactory(proxy.NewDefaultFactory(opencensus.BackendFactory(bf), logger)),
		Middlewares:    []gin.HandlerFunc{},
		Logger:         logger,
		HandlerFactory: opencensusgin.New(krakendgin.EndpointHandler),
	})

	// start the engine
	routerFactory.NewWithContext(ctx).Run(serviceConfig)
}
