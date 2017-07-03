package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/judwhite/go-svc/svc"
	options "github.com/mreiferson/go-options"

	app "github.com/molon/go_app_base/connector"
	"github.com/molon/go_app_base/internal/logger"
	"github.com/molon/go_app_base/internal/version"
)

type config map[string]interface{}

type program struct {
	app *app.Connector
}

func appFlagSet(opts *app.Options) *flag.FlagSet {
	flagSet := flag.NewFlagSet(app.APP_NAME, flag.ExitOnError)

	flagSet.String("config", "", "path to config file")
	flagSet.Bool("version", false, "print version string")

	flagSet.String("log-level", "info", "set log verbosity: debug, info, warn, error, or fatal")

	return flagSet
}

func newStdLogger(opts *app.Options) (*logger.Logger, error) {
	logLevel, err := logger.ParseLogLevel(opts.LogLevel)
	if err != nil {
		return nil, err
	}
	colors := true
	// Check to see if stderr is being redirected and if so turn off color
	// Also turn off colors if we're running on Windows where os.Stderr.Stat() returns an invalid handle-error
	stat, err := os.Stderr.Stat()
	if err != nil || (stat.Mode()&os.ModeCharDevice) == 0 {
		colors = false
	}
	return logger.NewStdLogger(true, colors, true, logLevel), nil
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *program) Start() error {
	opts := app.NewOptions()

	flagSet := appFlagSet(opts)
	flagSet.Parse(os.Args[1:])

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(version.String(app.APP_NAME))
		os.Exit(0)
	}

	var cfg map[string]interface{}
	configFile := flagSet.Lookup("config").Value.String()
	if configFile != "" {
		_, err := toml.DecodeFile(configFile, &cfg)
		if err != nil {
			log.Fatalf("ERROR: failed to load config file %s - %s", configFile, err.Error())
		}
	}

	options.Resolve(opts, flagSet, cfg)

	logger, err := newStdLogger(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	daemon := app.New(opts, logger)
	daemon.Main()
	p.app = daemon

	return nil
}

func (p *program) Stop() error {
	if p.app != nil {
		p.app.Exit()
	}
	return nil
}
