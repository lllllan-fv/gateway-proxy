package test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/lllllan-fv/gateway-proxy/public/conf"
)

func TestInitModule(t *testing.T) {
	c := qt.New(t)
	c.Helper()

	var config struct {
		App   conf.App
		MySQL conf.MySQL
		Redis conf.Redis
	}
	conf.InitModule("./config.toml", &config)
	{
		c.Assert(config.App.Env, qt.Equals, "production")
		c.Assert(config.App.Port, qt.Equals, uint16(30001))

		config.App.InitModule()
		c.Assert(config.App.Env, qt.Equals, conf.ProdEnv)

		c.Assert(config.MySQL.Addr, qt.Equals, "localhost:3307")
		c.Assert(config.MySQL.User, qt.Equals, "lllllan")
		c.Assert(config.MySQL.DBName, qt.Equals, "xinyoudui")
		c.Assert(config.MySQL.Password, qt.Equals, "123456")

		c.Assert(config.Redis.Addr, qt.Equals, "localhost:6379")
		c.Assert(config.Redis.Password, qt.Equals, "")
		c.Assert(config.Redis.Prefix, qt.Equals, "")
		c.Assert(config.Redis.DB, qt.Equals, uint(0))
	}
}
