package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	clearAWSRelatedEnv()

	envMap, err := godotenv.Read()

	if err != nil {
		log.Println("no .env file found, using system env")
	}

	keys := []string{
		"AWS_REGION",
		"AWS_PROFILE",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
	}

	for _, key := range keys {
		if val, ok := envMap[key]; ok {
			os.Setenv(key, val)
		}
	}
}

func clearAWSRelatedEnv() {
	os.Unsetenv("AWS_PROFILE")
	os.Unsetenv("AWS_REGION")
	os.Unsetenv("AWS_DEFAULT_REGION")
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")
	os.Unsetenv("AWS_SESSION_TOKEN")
}
