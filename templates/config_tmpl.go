package templates

const Config_common_tmpl = `
[common]
connstr = "{{.Connstr}}"         #以太坊连接地址
keydir = "{{.Keydir}}"           #以太坊私钥目录
ContractAddr = "{{.ContractAddr}}"  #合约地址，部署后得到
Buildpath  =  "{{.Buildpath}}"   #编译目标目录，默认为build
Codepath  =  "{{.Codepath}}"     #编译后的代码文件名
TestCodepath  =  "{{.TestCodepath}}"  #自动生成的测试用代码文件名 
TestConfigpath  =  "{{.TestConfigpath}}" #自动生成的配置文件名
`
const Config_deploy_tmpl = `
[Deploy]
func =  "{{.DeployFuncName}}"  #部署函数入口
fromaddr  =  "{{.FromAddr}}"   #部署用账户
pass  = "{{.Pass}}"            #部署账户密码
TestAddr = "{{.TestAddr}}"     #测试用账户
TestPass = "{{.TestPass}}"     #测试用账户密码
`
const Config_go_tmpl = `
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
	Common CommonConfig
	Deploy DeployConfig
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
	TestAddr string
	TestPass string
}

var Config *ServerConfig //引用配置文件结构

func init() {
	fmt.Println("call build config.init")
	Config = GetConfig()
	fmt.Println("call build config.init ok")
}

func GetConfig() (config *ServerConfig) {
	config = &ServerConfig{}
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	return config
}
`
