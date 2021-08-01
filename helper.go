package helper

import (
	"fmt"

	"github.com/santoshanand/helper/config"
	"github.com/santoshanand/helper/logger"
	"go.uber.org/zap"
)

var zlog *zap.SugaredLogger

// Initialize - initialie helper it will load config and initialize logger
func Initialize() error {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("configuration initialization error")
		return err
	}
	zlog = logger.InitLogger(false)
	return nil
}

// GetConfigValue - get value for given key
func GetConfigValue(key string) string {
	if key == "" {
		return ""
	}
	return config.GetValue(key)
}

// Log - print log info on console.
func Log(args ...interface{}) {
	zlog.Info(args...)
}

// Logf - print formatted log info on console.
func Logf(template string, args ...interface{}) {
	zlog.Infof(template, args...)
}
