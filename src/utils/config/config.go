package config

import (
	"github.com/spf13/viper"
)

func init() {
	// Viper will check for an environment variable any time a viper.Get request is made.
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
}

func Get(key string) interface{} {
	return viper.Get(key)
}
