package main

import (
	_ "fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/fv0008/AWS_Russia/server/Gamecontrollers"
	"github.com/fv0008/AWS_Russia/server/Global"
	"github.com/fv0008/AWS_Russia/server/Global/Game"
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
	go Game.GameRussia()
	time.Sleep(2 * time.Second)
	beego.Router("/", &MainController{},"get:Single")
	beego.Router("/watch", &MainController{},"get:Watch")
	beego.Router("/game", &MainController{},"get:Game")
	beego.Router("/IM/", &Gamecontrollers.AppController{})
	beego.Router("/IM/join", &Gamecontrollers.AppController{}, "post:Join")

	// WebSocket.
	beego.Router("/IM/ws", &Gamecontrollers.WebSocketController{})
	beego.Router("/IM/ws/socket", &Gamecontrollers.WebSocketController{}, "get:Socket")

	beego.Run()

}