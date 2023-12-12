package gateway

import (
	"cart-manager/infra/httpsrv"
	"cart-manager/registry"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r registry.ServiceRegistry) (*http.Server, error) {

	engine := gin.Default()
	// gin.SetMode(gin.ReleaseMode)

	address, err := httpsrv.GetListenAddressEnv()
	if err != nil {
		return nil, err
	}

	server := &http.Server{Addr: address, Handler: engine}
	return server, nil
}
