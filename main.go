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

	cfg := new(config.Config)
	err = rtmp.ListenAndServe(cfg)
	if err != nil {
		panic(err)
	}

	select {}

}
