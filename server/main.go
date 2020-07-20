package server

import (
	_ "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping("game", c.game)
	c.Mapping("watch", c.watch)
	c.Mapping("single", c.single)
}
func (this *MainController) game() {
	this.TplName = "h5Russia_client.html" // version 1.6 use this.TplName = "index.tpl"
}
func (this *MainController) watch() {
	this.TplName = "h5Russia_server.html" // version 1.6 use this.TplName = "index.tpl"
}
func (this *MainController) single() {
	this.TplName = "h5Russia_single.html" // version 1.6 use this.TplName = "index.tpl"
}
func main() {
	beego.Router("/", &MainController{},"get:single")
	beego.Router("/watch", &MainController{},"get:watch")
	beego.Router("/game", &MainController{},"get:game")
	return
}