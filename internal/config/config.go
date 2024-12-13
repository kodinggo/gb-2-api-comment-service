package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() {
	godotenv.Load()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func ENV() string {
	return viper.GetString("env")
}

func Port() string {
	return viper.GetString("port")
}

func MySQLDBHost() string {
	return viper.GetString("mysql.dbhost")
}

func MySQLDBUser() string {
	return viper.GetString("mysql.dbuser")
}

func MySQLDBPass() string {
	return viper.GetString("mysql.dbpass")
}

func MySQLDBName() string {
	return viper.GetString("mysql.dbname")
}
