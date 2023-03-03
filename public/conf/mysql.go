package conf

type MySQL struct {
	Addr            string `mapstructure:"addr"     default:"localhost:3306"`
	User            string `mapstructure:"user"     default:"root"`
	DBName          string `mapstructure:"db_name"  default:"mysql"`
	Password        string `mapstructure:"password" default:"123456"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}
