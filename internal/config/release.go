//+build !dev

package config

type Config struct {
	DebugMode bool `env:"DEBUG" envDefault:"false"`
	Verbose   bool `env:"VERBOSE" envDefault:"false"`
}
