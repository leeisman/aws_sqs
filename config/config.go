package configs

import (
	"fmt"
	viper "github.com/spf13/viper"
	"os"
)

type Config struct {
	Mysql  Mysql  `mapstructure:"mysql",yaml:"mysql,omitempty"`
	AwsSQS AwsSQS `mapstructure:"aws_sqs",yaml:"aws_sqs,omitempty"`
}

type Mysql struct {
	IP       string `mapstructure:"ip",yaml:"ip,omitempty"`
	Username string `mapstructure:"username,yaml:"username,omitempty"`
	Password string `mapstructure:"password,yaml:"password,omitempty"`
	Dbname   string `mapstructure:"db_name,yaml:"db_name,omitempty"`
}

type AwsSQS struct {
	QURL string `mapstructure:"qURL",yaml:"qURL,omitempty"`
}

func NewConfig() *Config {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	runtimeViper := viper.New()

	runtimeViper.SetConfigName("app")  // name of config file (without extension)
	runtimeViper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	runtimeViper.AddConfigPath(pwd)
	var config Config

	err = runtimeViper.ReadInConfig() // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = runtimeViper.Unmarshal(&config)
	if err != nil {
		fmt.Println("parse config fatal")
		os.Exit(1)
	}
	return &config
}
