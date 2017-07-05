package stdlib

import (
	"path/filepath"
	"os"
	"os/exec"
	"strings"
)

var hostname string

func GetHostName()string {
	h,_ := os.Hostname()
	return h
}

func GetWorkDir()string {
	file,_ := exec.LookPath(os.Args[0])
	path,_ := filepath.Abs(file)
	dir,_ := filepath.Split(path)
	path = strings.Replace(dir,"/bin","",1)
	return path
}

func init(){
	hostname = GetHostName()
}