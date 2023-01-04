package cfg

import "time"

type Config struct {
	BindAddr    string        `toml:"bind_addr"`
	SecretKey   string        `toml:"secret_key"`
	JWTtime     time.Duration `toml:"jwt_time"`
	DatabaseURL string        `toml:"database_url"`
	RedisHost   string        `toml:"redis_host"`
	RedisPort   string        `toml:"redis_port"`
	RedisPass   string        `toml:"redis_pass"`
	RedisDb     int           `toml:"redis_db"`
	RedisExp    time.Duration `toml:"redis_exp"`
	RabbitURL   string        `toml:"rabbit_url"`
	MailerExp   time.Duration `toml:"mailer_exp"`
}

var config *Config

// returning new default config
func NewConfig() *Config {
	config = &Config{
		BindAddr:  ":8080",
		JWTtime:   60,
		SecretKey: "e5je5seiop34[0",
		RedisPort: ":6379",
		RedisPass: "",
		RedisDb:   0,
		RedisExp:  60,
		MailerExp: 60,
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
