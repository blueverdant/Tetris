package main

import (
	_ "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping("Game", c.Game)
	c.Mapping("Watch", c.Watch)
	c.Mapping("Single", c.Single)
}
func (this *MainController) Game() {
	this.TplName = "h5Russia_client.html" // version 1.6 use this.TplName = "index.tpl"
}
func (this *MainController) Watch() {
	this.TplName = "h5Russia_server.html" // version 1.6 use this.TplName = "index.tpl"
}
func (this *MainController) Single() {
	this.TplName = "h5Russia_single.html" // version 1.6 use this.TplName = "index.tpl"
}
func main() {
	beego.Router("/", &MainController{},"get:Single")
	beego.Router("/watch", &MainController{},"get:Watch")
	beego.Router("/game", &MainController{},"get:Game")
	beego.Run()
}