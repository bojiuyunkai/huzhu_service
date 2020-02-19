package logs

import (
    "github.com/astaxie/beego/logs"
)
func init(){
	logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}