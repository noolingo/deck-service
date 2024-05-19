package domain

import "time"

type Config struct {
	AppName     string     `yaml:"appname" env-default:"test"`
	GrpcServer  GrpcServer `yaml:"grpcserver" env-prefix:"DECK_SERVICE_"`
	Mysql       Mysql      `yaml:"mysql" env-prefix:"DECK_SERVICE_"`
	CardService string     `yaml:"cardservice" env:"GRPC_CLIENTS_CARDSERVICE"`
}

type GrpcServer struct {
	Host string `yaml:"host" env-default:"0.0.0.0" env:"HOST"`
	Port string `yaml:"port" env-default:"9002" env:"PORT"`
}

type Mysql struct {
	DSN             string        `yaml:"dsn" env:"MYSQL_DSN"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env-default:"5m"`
	MaxOpenConns    int           `yaml:"max_open_conns" env-default:"10"`
	MaxIdleConns    int           `yaml:"max_idle_conns" env-default:"10"`
}
