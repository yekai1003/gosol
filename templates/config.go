package templates

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
	Common CommonConfig
	Deploy DeployConfig
}

type CommonConfig struct {
	Connstr        string
	Keydir         string
	ContractAddr   string
	Buildpath      string
	Codepath       string
	TestCodepath   string
	TestConfigpath string
}
type DeployConfig struct {
	ContractName string
	CallFunc     string
	FromAddr     string
	Pass         string
	AbiFile      string
}

var Config *ServerConfig //引用配置文件结构

func init() {
	Config = GetConfig()
	fmt.Println("templates init ok")
}

func GetConfig() (config *ServerConfig) {
	config = &ServerConfig{}
	if _, err := toml.DecodeFile("templates/templates.toml", &config); err != nil {
		panic(err)
	}

	return config
}
