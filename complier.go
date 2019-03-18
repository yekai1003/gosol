package main

import (
	"fmt"
	_ "io/ioutil"
	"os/exec"
	"strings"
)

func Compiler(solName, solPath, targetPath string) (string, error) {
	abigenpath, err := exec.LookPath("abigen")
	if err != nil {
		fmt.Println("failed to get abigen", err)
		return "", err
	}
	goName := strings.Replace(solName, ".sol", ".go", 4)
	fmt.Println(abigenpath)
	cmd := exec.Command(abigenpath, "-sol", solPath+"/"+solName, "-pkg", targetPath, "-out", targetPath+"/"+goName)
	err = cmd.Run()
	// reader, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println("failed to get stdoutpipe", err)
	// 	return "", err
	// }
	// //defer reader.Close()
	// cmd.Start()
	// data, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	fmt.Println("failed to read stdout pipe", err)
	// 	return "", err
	// }
	// cmd.Wait()
	fmt.Println("run complie go ok!!", err)
	return targetPath + "/" + goName, err
}

func BuildAbi(goCodeName string) error {
	abiName := strings.Replace(goCodeName, ".go", ".abi", 3)
	cmd := exec.Command("/bin/bash", "abi.sh", goCodeName, abiName)
	err := cmd.Run()
	fmt.Println("run BuildAbi ok!!", err)
	return nil
}
