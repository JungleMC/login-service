//+build dev

package config

type Config struct {
	DebugMode bool `env:"DEBUG" envDefault:"true"`
	Verbose   bool `env:"VERBOSE" envDefault:"true"`
}
