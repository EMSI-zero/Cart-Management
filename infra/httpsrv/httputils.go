package httpsrv

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var EnvListenAddress = "SRV_LISTEN_ADDRESS"

func StartServer(httpserver *http.Server) (err error) {

	go func() {
		logrus.Infof("http server listening on %s", httpserver.Addr)
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
