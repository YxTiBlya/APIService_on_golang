package cfg

import (
	"time"
)

type Config struct {
	BindAddr          string        `toml:"bind_addr"`
	SecretKey         string        `toml:"secret_key"`
	JWTtime           time.Duration `toml:"jwt_time"`
	Postgres_user     string        `toml:"postgres_user"`
	Postgres_password string        `toml:"postgres_password"`
	Postgres_host     string        `toml:"postgres_host"`
	Postgres_port     string        `toml:"postgres_port"`
	Postgres_dbname   string        `toml:"postgres_dbname"`
	Postgres_ssl      string        `toml:"postgres_ssl"`
	RedisHost         string        `toml:"redis_host"`
	RedisPort         string        `toml:"redis_port"`
	RedisPass         string        `toml:"redis_pass"`
	RedisDb           int           `toml:"redis_db"`
	RedisExp          time.Duration `toml:"redis_exp"`
	RabbitURL         string        `toml:"rabbit_url"`
	MailerExp         time.Duration `toml:"mailer_exp"`
}

var config *Config

// returning new default config
func NewConfig() *Config {
	config = &Config{
		BindAddr:          ":8080",
		SecretKey:         "e5je5seiop34[0",
		JWTtime:           60,
		Postgres_user:     "postgres",
		Postgres_password: "password",
		Postgres_port:     "5432",
		RedisPort:         ":6379",
		RedisPass:         "",
		RedisDb:           0,
		RedisExp:          60,
		MailerExp:         60,
	}
	return config
}

// return config struct
func GetConfig() *Config {
	return config
}

// change config data
func ChangeConfig(c *Config) {
	config = c
}
