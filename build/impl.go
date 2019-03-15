package main

import (
	"fmt"
	"gosol/templates"
	"os"
	"text/template"
)

type MainSolImpl struct {
	Connstr string
	Keydir  string
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
}
