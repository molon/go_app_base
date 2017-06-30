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
	flagSet := flag.NewFlagSet(app.AppName, flag.ExitOnError)

	flagSet.String("config", "", "path to config file")
	flagSet.Bool("version", false, "print version string")

	flagSet.String("log-level", "info", "set log verbosity: debug, info, warn, error, or fatal")

	return flagSet
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
		fmt.Println(version.String(app.AppName))
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

	logLevel, err := logger.ParseLogLevel(opts.LogLevel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger := logger.NewStdLogger(true, true, true, logLevel)

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
