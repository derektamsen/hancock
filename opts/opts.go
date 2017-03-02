package opts

import (
	"log"

	"github.com/spf13/viper"
)

// GetConfigString sets default values for missing config options, reads the ENV var, and returns
//  value for the option.
// Order:
// 1. ENV variable
// 2. value from config.<yaml, json, toml>
// 3. Hancock default
func GetConfigString(opt string, def string) string {
	viper.SetDefault(opt, def)
	viper.BindEnv(opt)
	return viper.GetString(opt)
}

// GetConfigInt sets default values for missing config options, reads the ENV var, and returns
//  value for the option.
// Order:
// 1. ENV variable
// 2. value from config.<yaml, json, toml>
// 3. Hancock default
func GetConfigInt(opt string, def int) int {
	viper.SetDefault(opt, def)
	viper.BindEnv(opt)
	return viper.GetInt(opt)
}

// SetConfigLoad reads the configuration file.
// Can be in the format of JSON, TOML, YAML, HCL, and Java Properties files
// reads `./config.<format>`
func SetConfigLoad(envprefix string, configFile string) {
	viper.SetEnvPrefix(envprefix)
	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print(err, " Continuing with default and ENV options")
	}
}
