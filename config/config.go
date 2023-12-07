/*

 */

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config holds the config data object with the actual data. Methods written over Config object give
// access to actual data.
// This is to make configurations read-only and also allow easy parsing of configuration files.
type Config struct {
	config configData
}

type configData struct {
	// http
	HTTPPort int    `mapstructure:"http_port"`
	GinMode  string `mapstructure:"gin_mode"`

	// log
	LogLevel        logrus.Level `mapstructure:"log_level"`
	LogFileLocation string       `mapstructure:"log_file_location"`

	// database
	DBType           string `mapstructure:"db_type"`
	DBConnectionPath string `mapstructure:"db_path"`

	//AWS Config
	AWSConfig AwsConfig `mapstructure:"awsconfig"`

	//TxProcessorConfig TxProcessorConfig `mapstructure:"tx_processor_config"`
}

type AwsConfig struct {
	CognitoRegion     string `mapstructure:"cognito_region"`
	CognitoUserPoolID string `mapstructure:"cognito_user_pool_id"`
	S3BucketName      string `mapstructure:"s3_bucket_name"`
	SQSNameTx         string `mapstructure:"sqs_queue_name_tx"`
}

func NewConfig() (*Config, error) {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")

	// where to look for
	viper.AddConfigPath("/etc/bitspawn/api/") // production config path
	viper.AddConfigPath("./config")           // dev config path
	viper.AddConfigPath("../config")          // dev config path

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	// Override config file based on ENV variables e.g. DB_TYPE=postgres
	viper.AutomaticEnv()

	config := configData{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &Config{config: config}, nil
}

func (c *Config) HTTPPort() int {
	return c.config.HTTPPort
}

func (c *Config) GinMode() string {
	return c.config.GinMode
}

func (c *Config) LogFileLocation() string {
	return c.config.LogFileLocation
}

func (c *Config) LogLevel() logrus.Level {
	return c.config.LogLevel
}

func (c *Config) DBCredentials() []string {
	return []string{
		c.config.DBType,
		c.config.DBConnectionPath,
	}
}

func (c *Config) AwsConfig() AwsConfig {
	return c.config.AWSConfig
}

//type TxProcessorConfig struct {
//	TaskLogFile       string `mapstructure:"task_log_file"`
//	// # of tasks in the pool
//	TaskPoolCount int `mapstructure:"task_pool_count"`
//	// # of functions in task routine
//	TaskSize  int `mapstructure:"task_size"`
//}
