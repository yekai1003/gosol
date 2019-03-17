package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var (
	Version   = "1.0.0"
	Commit    = "2019-03-15"
	BuildTime = "2019-03-15"
)

type ServerConfig struct {
	Common         CommonConfig
	Deploy         DeployConfig
	FuncConfs      []FuncConfig
	NoGasFuncConfs []NoGasFuncConfig
}

type CommonConfig struct {
	Connstr      string
	Keydir       string
	ContractAddr string
	Codepath     string
	TestCodepath string
}
type DeployConfig struct {
	Func     string
	FromAddr string
	Pass     string
	Params   string
}
type FuncConfig struct {
	Func           string
	IncontractName string
	Value          int
	Funcparams     string
}
type NoGasFuncConfig struct {
	Func           string
	IncontractName string
	RetParams      string
	Funcparams     string
}

var Config *ServerConfig //引用配置文件结构

func init() {
	fmt.Println("call config.init")
	Config = GetConfig()
}

func GetConfig() (config *ServerConfig) {
	config = &ServerConfig{}
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	return config
}
