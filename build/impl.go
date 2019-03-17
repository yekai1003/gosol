package main

import (
	"fmt"
	"gosol/templates"
	"os"
	"os/exec"
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

func ImplRunCode() {
	outfile, err := os.OpenFile(Config.Common.Codepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("failed to create  file:", Config.Common.Codepath, err)
		return
	}
	defer outfile.Close()
	tmpl_main, err := template.New("main").Parse(templates.Main_tmpl)
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
	tmpl_deploy, err := template.New("deploy").Parse(templates.Deploy_sol_tmpl)
	if err != nil {
		fmt.Println("failed to create  template tmpl_deploy:", err)
		return
	}
	err = tmpl_deploy.Execute(outfile, Config.Deploy)
	if err != nil {
		fmt.Println("failed to impl  template:", err, Config.Deploy)
		return
	}
	//完成要测试函数的代码组织 - 需要签名
	gasfunc_impl, err := template.New("testfunc").Parse(templates.Call_func_tmpl)
	for _, v := range Config.FuncConfs {
		err = gasfunc_impl.Execute(outfile, v)
		if err != nil {
			fmt.Println("failed to impl gasfunc_impl", err)
			break
		}
	}
	//完成要测试函数的代码组织 - 不需要签名
	nogasfunc_impl, err := template.New("testnogasfunc").Parse(templates.Call_nogas_func_tmpl)
	for _, v := range Config.NoGasFuncConfs {
		err = nogasfunc_impl.Execute(outfile, v)
		if err != nil {
			fmt.Println("failed to impl nogasfunc_impl", err)
			break
		}
	}

	//生成测试部分的代码

	outfile2, err := os.OpenFile(Config.Common.TestCodepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("failed to create  file:", Config.Common.Codepath, err)
		return
	}
	defer outfile2.Close()
	outfile2.WriteString(templates.Test_func_tmpl)

	str1 := fmt.Sprintf(templates.Test_func_run_tmpl, Config.Deploy.Func)
	outfile2.WriteString(str1)
	num := 3
	for _, v := range Config.FuncConfs {
		str := fmt.Sprintf(templates.Test_func_run_tmpl2, num, v.Func, Config.Deploy.FromAddr, Config.Deploy.Pass)
		outfile2.WriteString(str)
		num++
	}
	for _, v := range Config.NoGasFuncConfs {
		str := fmt.Sprintf(templates.Test_func_run_tmpl3, num, v.Func)
		outfile2.WriteString(str)
		num++
	}
	outfile2.WriteString("\n}")
	cmd := exec.Command("go", "build", "-i")
	cmd.Run()
}
