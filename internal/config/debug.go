//+build dev

package config

type Config struct {
	DebugMode bool `env:"DEBUG" envDefault:"true"`
	Verbose   bool `env:"VERBOSE" envDefault:"true"`

	// Redis
	RedisHost string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort int `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDatabase int`env:"REDIS_DB" envDefault:"0"`
}
