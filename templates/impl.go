package templates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	_ "os/exec"
	"strings"
	"text/template"
)

type MainSolImpl struct {
	Connstr string
	Keydir  string
}

type TestFuncImpl struct {
	DeployName string
	Keydir     string
}
type FuncConfigData struct {
	FuncName    string
	Constructor string
	CallFunc    string
	RetFunc     string
}

func ImplRunCode() {
	fmt.Println("ImplRunCode() begin")
	//需要读取abi文件自动填写，不再根据配置文件

	data, err := readAbiFile(Config.Deploy.AbiFile)
	if err != nil {
		fmt.Println("failed to read abi file ", err)
		return
	}
	var abis []ContractAbi

	//fmt.Println(data)

	err = json.Unmarshal([]byte(data), &abis)
	if err != nil {
		fmt.Println("failed to unmarshal json", err)
		return
	}

	outfile, err := os.OpenFile(Config.Common.Buildpath+"/"+Config.Common.Codepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("failed to create  file:", Config.Common.Codepath, err)
		return
	}
	defer outfile.Close()
	tmpl_main, err := template.New("main").Parse(Main_tmpl)
	if err != nil {
		fmt.Println("failed to create  template:", err)
		return
	}
	main_teml_data := MainSolImpl{Config.Common.Connstr, Config.Common.Keydir}
	err = tmpl_main.Execute(outfile, main_teml_data)
	if err != nil {
		fmt.Println("failed to impl  template:", err, main_teml_data)
		return
	}
	//上面完成第一部分

	//完成要测试函数的代码组织 - 需要签名
	gasfunc_impl, err := template.New("testfunc").Parse(Call_func_tmpl)
	if err != nil {
		fmt.Println("failed to parse Call_func_tmpl", err, Call_func_tmpl)
		return
	}
	//完成要测试函数的代码组织 - 不需要签名
	nogasfunc_impl, err := template.New("testnogasfunc").Parse(Call_nogas_func_tmpl)
	if err != nil {
		fmt.Println("failed to parse Call_nogas_func_tmpl", err, Call_nogas_func_tmpl)
		return
	}

	//生成测试部分的代码

	outfile2, err := os.OpenFile(Config.Common.Buildpath+"/"+Config.Common.TestCodepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("failed to create  file:", Config.Common.TestCodepath, err)
		return
	}
	defer outfile2.Close()
	outfile2.WriteString(Test_func_tmpl)

	str1 := fmt.Sprintf(Test_func_run_tmpl, "Deploy"+Config.Deploy.ContractName)
	outfile2.WriteString(str1)
	num2 := 3

	var funconfData FuncConfigData
	funconfData.Constructor = "New" + Config.Deploy.ContractName
	strFuncInit := ""
	for _, v := range abis {
		fmt.Println(v)
		if v.Type == "constructor" {
			//获取部署函数所要的参数
			if len(v.Inputs) > 0 {
				//有输入参数
				Config.Deploy.CallFunc = "Deploy" + Config.Deploy.ContractName + "(auth,testclient"
				for _, vv := range v.Inputs {
					if vv.Type == "address" {
						//common.HexToAddress
						Config.Deploy.CallFunc += ",common.HexToAddress(Config.Deploy.TestAddr)"
					} else if vv.Type == "uint256" {
						Config.Deploy.CallFunc += ",big.NewInt(10000)"
					} else if vv.Type == "string" {
						Config.Deploy.CallFunc += ",\"yekai\""
					}

				}
				Config.Deploy.CallFunc += ")"

			} else {
				Config.Deploy.CallFunc = "Deploy" + Config.Deploy.ContractName + "(auth,testclient)"
			}
			continue
		}
		strFuncInit += fmt.Sprintf("\t{\"%s\",%d},\n", v.Name, num2)
		if len(v.Outputs) > 0 {
			//no gas func
			funconfData.FuncName = Capitalize(v.Name)
			fmt.Println(funconfData.FuncName, "no gas ")
			num := 0
			retParams := ""
			for _, _ = range v.Outputs {
				str := fmt.Sprintf("data%d,", num)
				num++
				retParams += str
			}
			retParams += "err"
			funconfData.RetFunc = retParams
			if len(v.Inputs) > 0 {
				//有输入参数
				funconfData.CallFunc = funconfData.FuncName + "(nil"
				for _, vv := range v.Inputs {
					if vv.Type == "address" {
						//common.HexToAddress
						funconfData.CallFunc += ",common.HexToAddress(Config.Deploy.TestAddr)"
					} else if strings.Index("uint", vv.Type) > 0 {
						funconfData.CallFunc += ",big.NewInt(10000)"
					} else if vv.Type == "string" {
						funconfData.CallFunc += ",\"testXXX\""
					}

				}
				funconfData.CallFunc += ")"

			} else {
				funconfData.CallFunc = funconfData.FuncName + "(nil)"
			}

			err = nogasfunc_impl.Execute(outfile, funconfData)
			if err != nil {
				fmt.Println("failed to exec tmpl", err)
				break
			}
			//调用测试部分代码生成
			str := fmt.Sprintf(Test_func_run_tmpl3, num2, funconfData.FuncName)
			outfile2.WriteString(str)
			num2++
			//fmt.Println(funconfData)
		} else {
			// gas func
			funconfData.FuncName = Capitalize(v.Name)
			//fmt.Println(funconfData.FuncName, "have gas ")
			//funconfData.Constructor = "New" + Config.Deploy.ContractName
			if len(v.Inputs) > 0 {
				//有输入参数
				funconfData.CallFunc = funconfData.FuncName + "(auth"
				for _, vv := range v.Inputs {
					if vv.Type == "address" {
						//common.HexToAddress
						funconfData.CallFunc += ",common.HexToAddress(Config.Deploy.TestAddr)"
					} else if vv.Type == "uint256" {
						funconfData.CallFunc += ",big.NewInt(10000)"
					} else if vv.Type == "string" {
						funconfData.CallFunc += ",\"testXXX\""
					}

				}
				funconfData.CallFunc += ")"

			} else {
				funconfData.CallFunc = funconfData.FuncName + "(auth)"
			}
			//fmt.Println(funconfData)
			err = gasfunc_impl.Execute(outfile, funconfData)
			if err != nil {
				fmt.Println("failed to exec tmpl", err)
				break
			}
			//调用测试部分代码生成

			str := fmt.Sprintf(Test_func_run_tmpl2, num2, funconfData.FuncName, "Config.Deploy.TestAddr", "Config.Deploy.Pass")
			outfile2.WriteString(str)
			num2++
		}

	}

	outfile2.WriteString("\n}")

	strFuncInit = fmt.Sprintf(Test_funcinit_tmpl, strFuncInit)
	outfile2.WriteString(strFuncInit)

	//部署函数生成
	tmpl_deploy, err := template.New("deploy").Parse(Deploy_sol_tmpl)
	if err != nil {
		fmt.Println("failed to create  template tmpl_deploy:", err)
		return
	}
	fmt.Println("\n-------------\n", Config.Deploy, "\n-------------")

	err = tmpl_deploy.Execute(outfile, Config.Deploy)
	if err != nil {
		fmt.Println("failed to impl  template:", err, Config.Deploy)
		return
	}
	// cmd := exec.Command("go", "build", "-i")
	// cmd.Run()
	fmt.Println("ImplRunCode() end")
}

type ConfigDeployData struct {
	DeployFuncName string
	FromAddr       string
	Pass           string
	TestAddr       string
	TestPass       string
}

type CallParams struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ContractAbi struct {
	Constant        bool         `json:"constant"`
	Inputs          []CallParams `json:"inputs"`
	Name            string       `json:"name"`
	Outputs         []CallParams `json:"outputs"`
	Payable         bool         `json:"payable"`
	StateMutability string       `json:"stateMutability"`
	Type            string       `json:"type"`
}

//生成配置文件
func ImplConfigCode() {
	fmt.Println("deploy config begin......")
	outfile, err := os.OpenFile(Config.Common.Buildpath+"/"+Config.Common.TestConfigpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("failed to create  file:", Config.Common.TestConfigpath, err)
		return
	}
	defer outfile.Close()
	//
	common_config_tmpl, err := template.New("common").Parse(Config_common_tmpl)
	if err != nil {
		fmt.Println("failed to create  template:", err)
		return
	}
	//main_teml_data := MainSolImpl{Config.Common.Connstr, Config.Common.Keydir}
	err = common_config_tmpl.Execute(outfile, Config.Common)
	if err != nil {
		fmt.Println("failed to impl  template:", err, Config.Common)
		return
	}

	deploy_config_tmpl, err := template.New("deploy").Parse(Config_deploy_tmpl)
	if err != nil {
		fmt.Println("failed to create  template:", err)
		return
	}
	var deploy_config_data ConfigDeployData
	deploy_config_data.DeployFuncName = "Deploy" + Config.Deploy.ContractName
	deploy_config_data.FromAddr = Config.Deploy.FromAddr
	deploy_config_data.Pass = Config.Deploy.Pass
	deploy_config_data.TestAddr = Config.Deploy.FromAddr
	deploy_config_data.TestPass = Config.Deploy.Pass
	//main_teml_data := MainSolImpl{Config.Common.Connstr, Config.Common.Keydir}
	err = deploy_config_tmpl.Execute(outfile, deploy_config_data)
	if err != nil {
		fmt.Println("failed to impl deploy_config_tmpl ", err, deploy_config_data)
		return
	}
	fmt.Println("deploy config ok......")
	return

	//接下来需要读取ABI文件内容
	// data, err := readAbiFile(Config.Deploy.AbiFile)
	// if err != nil {
	// 	fmt.Println("failed to read abi file ", err)
	// 	return
	// }
	// var abis []ContractAbi

	// //fmt.Println(data)

	// err = json.Unmarshal([]byte(data), &abis)
	// if err != nil {
	// 	fmt.Println("failed to unmarshal json", err)
	// 	return
	// }
	// //fmt.Println(abis, err)
	// //实例化数据
	// func_config_tmpl, err := template.New("func").Parse(Config_func_tmpl)
	// if err != nil {
	// 	fmt.Println("failed to create  template:", err)
	// 	return
	// }
	// nogasfunc_config_tmpl, err := template.New("func2").Parse(Config_NoGasfunc_tmpl)
	// if err != nil {
	// 	fmt.Println("failed to create  template:", err)
	// 	return
	// }
	// var funconfData FuncConfigData
	// funconfData.Constructor = "New" + Config.Deploy.ContractName
	// //fmt.Println("funconfData====", funconfData)

	// for _, v := range abis {
	// 	fmt.Println(v)
	// 	if len(v.Outputs) > 0 {
	// 		//no gas func
	// 		funconfData.FuncName = Capitalize(v.Name)
	// 		fmt.Println(funconfData.FuncName, "no gas ")
	// 		num := 0
	// 		retParams := ""
	// 		for _, _ = range v.Outputs {
	// 			str := fmt.Sprintf("data%d,", num)
	// 			num++
	// 			retParams += str
	// 		}
	// 		retParams += "err"
	// 		funconfData.RetFunc = retParams
	// 		if len(v.Inputs) > 0 {
	// 			//有输入参数
	// 			funconfData.CallFunc = funconfData.FuncName + "(nil"
	// 			for _, vv := range v.Inputs {
	// 				if vv.Type == "address" {
	// 					//common.HexToAddress
	// 					funconfData.CallFunc += ",common.HexToAddress(Config.Deploy.TestAddr)"
	// 				} else if strings.Index("uint", vv.Type) > 0 {
	// 					funconfData.CallFunc += ",big.NewInt(10000)"
	// 				} else if vv.Type == "string" {
	// 					funconfData.CallFunc += ",\"testXXX\""
	// 				}

	// 			}
	// 			funconfData.CallFunc += ")"

	// 		} else {
	// 			funconfData.CallFunc = funconfData.FuncName + "(nil)"
	// 		}

	// 		err = nogasfunc_config_tmpl.Execute(outfile, funconfData)
	// 		if err != nil {
	// 			fmt.Println("failed to exec tmpl", err)
	// 			break
	// 		}
	// 		//fmt.Println(funconfData)
	// 	} else {
	// 		// gas func
	// 		funconfData.FuncName = Capitalize(v.Name)
	// 		//fmt.Println(funconfData.FuncName, "have gas ")
	// 		//funconfData.Constructor = "New" + Config.Deploy.ContractName
	// 		if len(v.Inputs) > 0 {
	// 			//有输入参数
	// 			funconfData.CallFunc = funconfData.FuncName + "(auth"
	// 			for _, vv := range v.Inputs {
	// 				if vv.Type == "address" {
	// 					//common.HexToAddress
	// 					funconfData.CallFunc += ",common.HexToAddress(Config.Deploy.TestAddr)"
	// 				} else if strings.Index("uint", vv.Type) > 0 {
	// 					funconfData.CallFunc += ",big.NewInt(10000)"
	// 				} else if vv.Type == "string" {
	// 					funconfData.CallFunc += ",\"testXXX\""
	// 				}

	// 			}
	// 			funconfData.CallFunc += ")"

	// 		} else {
	// 			funconfData.CallFunc = funconfData.FuncName + "(auth)"
	// 		}
	// 		//fmt.Println(funconfData)
	// 		err = func_config_tmpl.Execute(outfile, funconfData)
	// 		if err != nil {
	// 			fmt.Println("failed to exec tmpl", err)
	// 			break
	// 		}
	// 	}

	// }

	fmt.Println("ImplConfigCode config.toml ok......")
}

func readAbiFile(Abifile string) (string, error) {
	data, err := ioutil.ReadFile(Abifile)
	if err != nil {
		fmt.Println("failed to read file ", err)
		return "", err
	}
	data1 := strings.Replace(string(data), "\\", "", -1)
	return data1, nil
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func BuildConfigGO() {
	outfile, err := os.OpenFile(Config.Common.Buildpath+"/config.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("failed to create  file:", Config.Common.Buildpath, err)
		return
	}
	defer outfile.Close()
	outfile.WriteString(Config_go_tmpl)
}

func BuildRun() {
	ImplConfigCode()
	BuildConfigGO()
	ImplRunCode()
}
