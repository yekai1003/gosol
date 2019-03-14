package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func Compiler(solName, solPath, targetPath string) error {
	abigenpath, err := exec.LookPath("abigen")
	if err != nil {
		fmt.Println("failed to get abigen", err)
		return err
	}
	goName := strings.Replace(solName, ".sol", ".go", 4)
	fmt.Println(abigenpath)
	cmd := exec.Command(abigenpath, "-sol", solPath+"/"+solName, "-pkg", targetPath, "-out", targetPath+"/"+goName)

	reader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("failed to get stdoutpipe", err)
		return err
	}
	defer reader.Close()
	cmd.Start()
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("failed to read stdout pipe", err)
		return err
	}
	cmd.Wait()
	fmt.Println("run ok:", string(data))
	return err
}
