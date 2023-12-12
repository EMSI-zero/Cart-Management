package httpserver

import (
	"cart-manager/gateway"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var EnvListenAddress = "SRV_LISTEN_ADDRESS"

func StartServer() error {
	engine := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	gateway.AddRoutes(engine)

	address, err := GetListenAddressEnv()
	if err != nil {
		return err
	}

	httpserver := &http.Server{Addr: address, Handler: engine}
	go func() {
		logrus.Info("http server listening on %s", address)
		if httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Panic(err)
		}
	}()

	srv := httpserver
	chInterrupt := make(chan os.Signal, 1)
	signal.Notify(chInterrupt, os.Interrupt)
	chTerm := make(chan os.Signal, 1)
	signal.Notify(chTerm, syscall.SIGTERM)
	select {
	case <-chInterrupt:
		logrus.Info("server interrupted")
		stopServer(srv)
	case <-chTerm:
		logrus.Info("server received SIGTERM")
	}
	logrus.Info("done")

	return nil
}

func stopServer(srv *http.Server) {
	partStopped := make(chan struct{}, 20)
	go func() {
		const timeout = 5 * time.Second
		time.Sleep(timeout)
		partStopped <- struct{}{}
	}()
	go func() {
		logrus.Info(context.Background(), "stopping http server...")
		srvError := srv.Close()
		if srvError != nil {
			logrus.Info(context.Background(), "error: http server termination failed: %v", srvError)
		} else {
			logrus.Info(context.Background(), "stopped http server")
		}
		partStopped <- struct{}{}
	}()
	<-partStopped
}

func GetListenAddressEnv() (string, error) {
	address := os.Getenv(EnvListenAddress)
	if address == "" {
		return "", fmt.Errorf("env %s must be specified", EnvListenAddress)
	}
	return address, nil
}
