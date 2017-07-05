package controllers
import (
    "myblog/lib/runner"
    "github.com/astaxie/beego"
)

type BaseControllers struct{
    beego.Controller
}

type Response struct {
    RetCode         int       `json:"retCode"`
    Description     string      `json:"description"`
    Content         interface{} `json:"content"`
}

func (this * BaseControllers)Response(retcode int,c interface{}){
    req := &Response{
        RetCode: retcode,
        Description: runner.RetDescription[retcode],
        Content: c,
    }

    if err,ok := c.(error);ok {
        req.Description = err.Error()
        req.Content = nil
    }

    this.Data["json"] = req
    this.ServeJSON()
}
