package config

import (
	"github.com/spf13/viper"
)

type Configurator interface {
	Load() error
}

type Config struct {
	Name, Extension, Path string
}

func New(cnf Config) Configurator {
	return &ViperConfig{name: cnf.Name, ext: cnf.Extension, path: cnf.Path, cnf: viper.New()}
}

func GetString(name string) string {
	return viperConf.getString(name)
}

func GetStrings(name string) []string {
	return viperConf.getStrings(name)
}

func GetInt(name string) int {
	return viperConf.getInt(name)
}

func GetInts(name string) []int {
	return viperConf.getInts(name)
}

func GetFloat64(name string) float64 {
	return viperConf.getFloat64(name)
}

func GetBool(name string) bool {
	return viperConf.getBool(name)
}
