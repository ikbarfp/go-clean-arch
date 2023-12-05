package config

import (
	"github.com/spf13/viper"
)

var viperConf *ViperConfig

type ViperConfig struct {
	name, ext, path string
	cnf             *viper.Viper
}

func (v *ViperConfig) Load() error {
	v.cnf.SetConfigName(v.name)
	v.cnf.SetConfigType(v.ext)
	v.cnf.AddConfigPath(v.path)

	// Read all config
	if err := v.cnf.ReadInConfig(); err != nil {
		return err
	}

	v.cnf.AutomaticEnv()
	v.cnf.WatchConfig()

	viperConf = v

	return nil
}

func (v *ViperConfig) getString(name string) string {
	return v.cnf.GetString(name)
}

func (v *ViperConfig) getStrings(name string) []string {
	return v.cnf.GetStringSlice(name)
}

func (v *ViperConfig) getInt(name string) int {
	return v.cnf.GetInt(name)
}

func (v *ViperConfig) getInts(name string) []int {
	return v.cnf.GetIntSlice(name)
}

func (v *ViperConfig) getFloat64(name string) float64 {
	return v.cnf.GetFloat64(name)
}

func (v *ViperConfig) getBool(name string) bool {
	return v.cnf.GetBool(name)
}
