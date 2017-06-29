package connector

import (
	"log"
	"os"
	"sync"

	"github.com/molon/go_app_base/internal/util"
	"github.com/molon/go_app_base/internal/version"
)

type Connector struct {
	sync.RWMutex
	waitGroup util.WaitGroupWrapper
	opts      *Options
}

func New(opts *Options) *Connector {
	if opts.Logger == nil {
		opts.Logger = log.New(os.Stderr, opts.LogPrefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	n := &Connector{
		opts: opts,
	}

	n.logf(LOG_INFO, version.String("connector"))
	return n
}

func (l *Connector) Main() {
	//TODO:需要将一些本身阻塞的玩意丢到waitGroup里
	// l.waitGroup.Wrap(func() {
	// 	http_api.Serve(httpListener, httpServer, "HTTP", l.logf)
	// })
}

func (l *Connector) Exit() {
	//TODO:需要再这里去做清理操作
	l.waitGroup.Wait()
}
