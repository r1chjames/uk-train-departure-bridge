package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetStringEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		if len(fallback) > 0 {
			return fallback
		} else {
			panic(fmt.Sprintf("Required application variable %s not defined", key))
		}
	}
	return value
}

func GetIntEnv(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		if fallback > 0 {
			return fallback
		} else {
			panic(fmt.Sprintf("Required application variable %s not defined", key))
		}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Error parsing integer environment variable: %s", value))
	}
	return intValue
}