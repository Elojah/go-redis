package redis

// Config is redis structure config.
type Config struct {
	Addr     string `json:"addr" env:"REDIS_ADDR" env-default:"0.0.0.0"`
	Password string `json:"password" env:"REDIS_PASSWORD"`
	DB       int    `json:"db" env:"REDIS_DB"`
}
