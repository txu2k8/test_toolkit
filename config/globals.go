package config

import (
	"os"
	"time"

	"github.com/jinzhu/configor"
)

// define global const values
const (
	// Log level limited
	logLevel = 3
)

var (
	// Root work dir
	RootDir = os.Getenv("PWD")
	TimeStr = time.Now().Format("20060102150405")
)

// Config application all configs
var Config = struct {
	Worker struct {
		MaxNum     int    `env:"MAX_WORKER_NUM" default:"15"`
		VIPTenants string `env:"VIP_TENANTS"`
		V2Style    string `env:"V2_STYLE" default:"true"`
		Merge      struct {
			Enabled  string `env:"MERGE_ENABLED" default:"true"`
			PoolSize int    `env:"MERGE_POOL_SIZE" default:"100"`
		}
	}
	Version string `env:"Version" default:"v1.0.0"`
}{}

func init() {
	if err := configor.Load(&Config); err != nil {
		panic(err)
	}
}
