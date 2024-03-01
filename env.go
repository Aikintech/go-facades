package gofacades

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	dotenv "github.com/joho/godotenv"
)

type Env struct{}

func LoadEnv(filenames ...string) {
	err := dotenv.Load(filenames...)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(".env file not found, using only system environment variables.")
		} else {
			fmt.Printf("error loading .env file: %s", err)
		}
	}
}

func GetEnv() *Env {
	return &Env{}
}

// GetString gets string type env from os.
func (e *Env) GetString(key string, defaultValue ...interface{}) string {
	return e.env(key, defaultValue...).(string)
}

// GetInt gets int type env from os.
func (e *Env) GetInt(key string, defaultValue ...interface{}) int {
	val := e.env(key, defaultValue...)
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
func (e *Env) GetBool(key string, defaultValue ...interface{}) bool {
	val := e.env(key, defaultValue...)
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
func (e *Env) GetStringSlice(key string, defaultValue ...interface{}) []string {
	val := e.env(key, defaultValue...)
	split := strings.Split(val.(string), ",")

	return split
}

func (e *Env) env(key string, defaultValue ...interface{}) interface{} {
	val := os.Getenv(key)
	if val == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return val
}
