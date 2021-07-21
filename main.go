package Tetris

import (
	_ "fmt"
	"time"

	"github.com/Jugendreisen/Tetris/server/Gamecontrollers"
	"github.com/Jugendreisen/Tetris/server/Global"
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
	this.Data["IsWebSocket"] = true
	this.TplName = "h5Russia_server.html" // version 1.6 use this.TplName = "index.tpl"
}
func (this *MainController) Watch() {
	this.Data["IsWebSocket"] = true
	this.TplName = "h5Russia_client.html" // version 1.6 use this.TplName = "index.tpl"
}
func (this *MainController) Single() {
	this.TplName = "h5Russia_single.html" // version 1.6 use this.TplName = "index.tpl"
}
func main() {
	Global.Init_Logs()
	time.Sleep(2 * time.Second)
	beego.Router("/tetris/", &MainController{},"get:Single")
	beego.Router("/tetris/watch", &MainController{},"get:Watch")
	beego.Router("/tetris/game", &MainController{},"get:Game")
	beego.Router("/tetris/IM/", &Gamecontrollers.AppController{})
	beego.Router("/tetris/IM/join", &Gamecontrollers.AppController{}, "post:Join")

	// WebSocket.
	beego.Router("/tetris/IM/ws", &Gamecontrollers.WebSocketController{})
	beego.Router("/tetris/IM/ws/socket", &Gamecontrollers.WebSocketController{}, "get:Socket")
	beego.SetStaticPath("/tetris/static", "static")
	beego.Run()

}