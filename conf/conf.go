package conf

import (
	"time"

	"github.com/spf13/viper"
)

const (
	ConfigFilePath = "./config.yaml"
)

var v = viper.New()

func init() {
	// default config
	v.SetConfigFile(ConfigFilePath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Get(key string) interface{} {
	return v.Get(key)
}

func GetBool(key string) bool {
	return v.GetBool(key)
}

func GetFloat64(key string) float64 {
	return v.GetFloat64(key)
}

func GetInt(key string) int {
	return v.GetInt(key)
}

func GetInt32(key string) int32 {
	return v.GetInt32(key)
}

func GetInt64(key string) int64 {
	return v.GetInt64(key)
}

func GetIntSlice(key string) []int {
	return v.GetIntSlice(key)
}

func GetStringSlice(key string) []string {
	return v.GetStringSlice(key)
}

func GetString(key string) string {
	return v.GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return v.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return v.GetStringMapString(key)
}

func GetTime(key string) time.Time {
	return v.GetTime(key)
}

func GetUint(key string) uint {
	return v.GetUint(key)
}

func GetUint32(key string) uint32 {
	return v.GetUint32(key)
}

func GetUint64(key string) uint64 {
	return v.GetUint64(key)
}
