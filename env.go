package gofacades

import (
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	dotenv "github.com/joho/godotenv"
)

type env struct{}

func LoadEnv() {
	err := dotenv.Load()
	if err != nil {
		color.Redln("Failed to load .env: Switching to os")
	}
}

func Env() *env {
	return &env{}
}

// GetString gets string type env from os.
func (e *env) GetString(key string, defaultValue ...interface{}) string {
	return e.getEnv(key, defaultValue...).(string)
}

// GetInt gets int type env from os.
func (e *env) GetInt(key string, defaultValue ...interface{}) int {
	val := e.getEnv(key, defaultValue...)
	switch v := val.(type) {
	case int:
		return v
	case string:
		intVal, err := strconv.Atoi(v)
		if err == nil {
			return intVal
		}
	}
	return 0
}

// GetBool gets bool type config from application.
func (e *env) GetBool(key string, defaultValue ...interface{}) bool {
	val := e.getEnv(key, defaultValue...)
	switch v := val.(type) {
	case bool:
		return v
	case string:
		boolVal, err := strconv.ParseBool(v)
		if err == nil {
			return boolVal
		}
	}
	return false
}

// GetStringSlice gets string from env, split it by comma and converts to string slice.
func (e *env) GetStringSlice(key string, defaultValue ...interface{}) []string {
	val := e.getEnv(key, defaultValue...)
	split := strings.Split(val.(string), ",")

	return split
}

func (e *env) getEnv(key string, defaultValue ...interface{}) interface{} {
	val := os.Getenv(key)
	if val == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return val
}
