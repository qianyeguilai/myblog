package controllers

import (
 //   "github.com/astaxie/beego"
)

type TestController struct {
    BaseControllers
}

// @router /test [get,post]
func (this * TestController)Test(){
    result := "only test"
    this.Response(0,result)
}


