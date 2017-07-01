package connector

import (
	"sync"

	"github.com/molon/go_app_base/internal/logger"
	"github.com/molon/go_app_base/internal/util"
	"github.com/molon/go_app_base/internal/version"
)

const AppName = "ixora_connector"

type Connector struct {
	sync.RWMutex
	logger    *logger.Logger
	waitGroup util.WaitGroupWrapper
	opts      *Options
}

func New(opts *Options, logger *logger.Logger) *Connector {
	return &Connector{
		logger: logger,
		opts:   opts,
	}
}

func (c *Connector) Main() {
	c.Infof("Starting %s", version.String(AppName))

	//TODO:需要将一些本身阻塞的玩意丢到waitGroup里
	// c.waitGroup.Wrap(func() {
	// 	http_api.Serve(httpListener, httpServer, "HTTP", l.logf)
	// })
}

func (c *Connector) Exit() {
	c.Infof("Server is exiting")

	//TODO:需要再这里去做清理操作
	c.waitGroup.Wait()

	c.Infof("Server has exited")
}
