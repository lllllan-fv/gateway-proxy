package conf

type Redis struct {
	Addr         string `mapstructure:"addr"     default:"localhost:6379"`
	Password     string `mapstructure:"password" default:""`
	Prefix       string `mapstructure:""         default:""`
	DB           uint   `mapstructure:"db"       default:"0"`
	ConnTimeout  int    `mapstructure:"conn_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}
