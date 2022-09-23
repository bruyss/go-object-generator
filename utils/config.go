package utils

import "github.com/spf13/viper"

func GetConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.SetDefault("gensettings.general.second-pulse", "iSecPulse")
	viper.SetDefault("gensettings.general.simulation", "iSimulation")
	viper.SetDefault("gensettings.general.todo-bit", "\"DB_Test\".TODO_BOOL")
	viper.SetDefault("gensettings.general.todo-real", "\"DB_Test\".TODO_REAL")
	viper.SetDefault("gensettings.general.wincc", "true")

	err := viper.ReadInConfig()
	if err != nil {
		Sugar.Fatalln(err)
	}
}
