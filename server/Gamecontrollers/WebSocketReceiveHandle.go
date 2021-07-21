package Gamecontrollers

import (
	"github.com/Jugendreisen/Tetris/server/Global"
	"github.com/gorilla/websocket"
	"net/http"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/IM/", 302)
		return
	}

	this.TplName = "websocket.html"
	this.Data["IsWebSocket"] = true
	this.Data["UserName"] = uname
}



// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Socket() {

	SocketId,_ := this.GetUint32("SocketId")

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// Upgrade from http request to WebSocket.
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		Global.Logger.Error("Cannot setup deivce WebSocket connection:", err)
		return
	}
	// Join chat room. 后续所有的通信都不会在走这里而是走到join函数里循环
	globaWebSocketListManager.SocketJoin(SocketId,ws)

}



