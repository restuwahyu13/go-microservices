package packages

import (
	"github.com/spf13/viper"
)

func NewViper() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	return err
}
