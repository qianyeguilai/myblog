package runner

import (
	"os"
	"path/filepath"
	"fmt"
	"github.com/astaxie/beego"
	"myblog/lib/mycache"
	"myblog/lib/stdlib"
)

type RunnerConfig struct {
	//project basic info
	AppName			string	`form:"-"`
	WorkDir			string  `form:"-"`
	PidFile	        string	`form:"-"`
	WorkNum			string	`form:"-"`
	ApiDeactivate	bool	`form:"-"`
	CallDeactivate	bool	`form:"-"`
}

func (c * RunnerConfig)Refresh()error {
	_,err := os.Stat(filepath.Join(c.WorkDir,"logs/api_activate"))
	if err == nil || os.IsExist(err) {
		if c.ApiDeactivate != false {
			c.ApiDeactivate = false
			beego.Info("api activate done")
		}
	} else {
		if c.ApiDeactivate != true {
			c.ApiDeactivate = true
			beego.Info("api deactivate done")
		}
	}
    return nil
}

func (c * RunnerConfig)Rinterval()int64 {
	return 1
}

var (
	Rc =&RunnerConfig{
        WorkDir: stdlib.GetWorkDir(),
        ApiDeactivate: true,
        CallDeactivate: true,
	}
)

func ParseParams(config * RunnerConfig)(err error) {
    return nil
}

func Init() {
	mycache.AddCache(Rc)
	beego.Info("blog project workdir :",Rc.WorkDir)

	Rc.AppName = beego.AppConfig.String("appname")
	Rc.PidFile = filepath.Join(Rc.WorkDir,"logs/blog.pid")
	beego.Info(fmt.Sprintf("this project appname is %s and PidFile in %s",Rc.AppName,Rc.PidFile))

	ParseParams(Rc)
	beego.Info(fmt.Sprintf("%+v",Rc))
}
