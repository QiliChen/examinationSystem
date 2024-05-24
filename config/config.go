package config

import (
	"examination_system/utils"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Name      string `yaml:"name"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

var AppConfig Config

func InitConfig() {
	// 读取 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("无法加载 .env 文件:", err)
	}

	// 读取 config.yaml 文件
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("无法打开配置文件:", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal("无法解析配置文件:", err)
	}

	// 解密密码
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("未找到 SECRET_KEY 环境变量")
	}

	decryptedPassword, err := utils.Decrypt(AppConfig.Database.Password, secretKey)
	if err != nil {
		log.Fatal("无法解密密码:", err)
	}
	AppConfig.Database.Password = decryptedPassword
}
