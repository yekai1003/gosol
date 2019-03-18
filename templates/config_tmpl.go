package templates

const Config_common_tmpl = `
[common]
connstr = "{{.Connstr}}"
keydir = "{{.Keydir}}"
ContractAddr = "{{.ContractAddr}}"
Buildpath  =  "{{.Buildpath}}"
Codepath  =  "{{.Codepath}}"
TestCodepath  =  "{{.TestCodepath}}"
TestConfigpath  =  "{{.TestConfigpath}}"
`
const Config_deploy_tmpl = `
[Deploy]
func =  "{{.DeployFuncName}}"
fromaddr  =  "{{.FromAddr}}"
pass  = "{{.Pass}}"
Params =  "{{.Params}}"
TestAddr = "{{.TestAddr}}"
TestPass = "{{.TestPass}}"
`
const Config_func_tmpl = `
[[funcConfs]]
func =  "{{.FuncName}}"
Constructor = "{{.Constructor}}"
value  = 0
funcparams  =  "{{.CallFunc}}"
`

const Config_NoGasfunc_tmpl = `
[[NoGasFuncConfs]]
func =  "{{.FuncName}}"
Constructor = "{{.Constructor}}"
funcparams  =  "{{.CallFunc}}"
retparams  =  "{{.RetFunc}}"
`
