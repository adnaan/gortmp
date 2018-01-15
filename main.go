package main

import (
	//"./mpegts"
	"github.com/adnaan/gortmp/rtmp"
	//"fmt"
	//"os"
	//"./avformat"
	"github.com/adnaan/gortmp/config"
)

func main() {
	var err error

	if err = InitAppConfig(); err != nil {
		return
	}

	l := ":1935"
	err = rtmp.ListenAndServe(l)
	if err != nil {
		panic(err)
	}

	select {}

}

func InitAppConfig() (err error) {
	cfg := new(config.Config)
	err = cfg.Init("app.conf")
	if err != nil {
		return
	}

	return
}
