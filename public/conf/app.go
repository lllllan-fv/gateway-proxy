package conf

const (
	ProdEnv = "prod"
	DecEnv = "dev"
)

type App struct {
	Env            string `mapstructure:"env"              default:"prod"` // environment
	Port           uint16 `mapstructure:"port"             default:"8080"` // Listening port
	ReadTimeout    uint16 `mapstructure:"read_timeout"     default:"10"`
	WriteTimeout   uint16 `mapstructure:"write_timeout"    default:"10"`
	MaxHeaderBytes uint16 `mapstructure:"max_header_bytes" default:"20"`
}

func (app *App) InitModule() {

	switch app.Env {
	case ProdEnv:
		app.Env = ProdEnv
	case DecEnv:
		app.Env = DecEnv
	default:
		app.Env = ProdEnv
	}
}

func (app *App) IsProdEnv() bool {
	return app.Env == ProdEnv
}

func (app *App) IsDevEnv() bool {
	return app.Env == DecEnv
}
