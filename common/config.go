package common

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type JwtConfig struct {
	SecretKey string `yaml:"secretKey"`
	Expire    int32  `yaml:"expire"`
}

type LogCOnfig struct {
	Level         zap.AtomicLevel       `yaml:"level"`
	Filepath      string                `yaml:"filepath"`
	EncoderConfig zapcore.EncoderConfig `yaml:"encoderConfig"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Id       string `yaml:"id"`
	Pw       string `yaml:"pw"`
	Database string `yaml:"database"`
}

type Config struct {
	Jwt   JwtConfig      `yaml:"jwt"`
	Log   LogCOnfig      `yaml:"log"`
	DB    DatabaseConfig `yaml:"db"`
	Redis DatabaseConfig `yaml:"redis"`
}

var config *Config

func init() {
	GetConfig()
}

func GetConfig() *Config {
	if config == nil {
		var mode string
		if mode = os.Getenv("APP_MODE"); mode == "" {
			mode = "local"
		}
		filename, _ := filepath.Abs(fmt.Sprintf("config/%s.yaml", mode))
		if b, err := os.ReadFile(filename); err != nil {
			panic(err)
		} else {
			if err2 := yaml.Unmarshal(b, &config); err2 != nil {
				panic(err2)
			}
		}

	}
	return config
}
