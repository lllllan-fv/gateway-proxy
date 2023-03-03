package conf

import (
	"flag"

	"github.com/lllllan-fv/gateway-proxy/public/conf"
)

type Config struct {
	App   conf.App
	MySQL conf.MySQL
	Redis conf.Redis
}

var config Config

func GetConfig() *Config {
	return &config
}

// go run main.go --config={path}
var path = flag.String("config", "./config", "input config file like ./dev/config.toml")

func init() {
	flag.Parse()

	conf.InitModule(*path, &config)

	config.App.InitModule()
}
