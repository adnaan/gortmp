package config

import "github.com/kelseyhightower/envconfig"

// Config contains ENV variables
type Config struct {
	//HLSFragment
	HLSFragment int64 `split_words:"true" default:"2"`
	// HLSWindow
	HLSWindow int `split_words:"true" default:"5"`
	//HLSPath
	HLSPath   string `split_words:"true" default:"/Volumes/ramdisk"`
	DebugMode bool   `split_words:"true" default:"false"`
	RtmpPort  int    `split_words:"true" default:"1935"`
}

// Load settles ENV variables into Config structure
func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
