package Gamecontrollers

import (
	"container/list"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/Jugendreisen/Tetris/server"
	"github.com/Jugendreisen/Tetris/server/Global"
	"github.com/Jugendreisen/Tetris/server/Global/Game"

	"github.com/gorilla/websocket"
)

var globaWebSocketListManager *WebSocketListController
//带用户信息的websocket
type SocketInfo struct {
	SocketId	uint32
	User     	server.IM_protocol_user
	Conn     	*websocket.Conn
}
type SocketId struct {
	SocketId	 uint32
}
type UnSocketId struct {
	SocketId	 uint32
}
type WebSocketListController struct {
	// Channel for new join users.
	SocketChan chan SocketInfo
	// Channel for exit users.
	UnSocketChan chan UnSocketId
	// Send events here to publish them.
	MsgList	chan(server.IM_protocol)
	// Long polling waiting list.
	ActiveSocketList *list.List
	beego.Controller
}



func (this *WebSocketListController)IsExistSocketById(SocketId uint32) (bool) {
	for sub := this.ActiveSocketList.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(SocketInfo).SocketId == SocketId {
			return true
		}
	}
	return false
}

func init(){
	globaWebSocketListManager =	&WebSocketListController{} ;
	globaWebSocketListManager.SocketChan = make(chan SocketInfo, 100)
	// Channel for exit users.
	globaWebSocketListManager.UnSocketChan = make(chan UnSocketId, 100)
	// Send events here to publish them.
	globaWebSocketListManager.MsgList = make(chan server.IM_protocol, 100)

	globaWebSocketListManager.ActiveSocketList = list.New()


	go globaWebSocketListManager.chatroom()
	go globaWebSocketListManager.NetRussia()
}

func (this *WebSocketListController)SocketLeave(SocketId uint32) {
	this.UnSocketChan <- UnSocketId{SocketId}
	Global.Logger.Info("Socket Leave:", SocketId)
}

func (this *WebSocketListController)SocketJoin(SocketId uint32,ws *websocket.Conn) {
	if this.IsExistSocketById(SocketId){
		for sub := this.ActiveSocketList.Front(); sub != nil; sub = sub.Next() {
			if sub.Value.(SocketInfo).SocketId == SocketId {
				this.ActiveSocketList.Remove(sub)
				break
			}
		}
	}

	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		NewSocketId := r.Uint32()
		if !this.IsExistSocketById(NewSocketId) {
			//这里就是整个用户存在的循环体积，先将用户放入订阅队列
			this.SocketChan <- SocketInfo{NewSocketId, server.IM_protocol_user{}, ws}
			//预定函数结尾让用户离开， 因为有可能强行kick，所以有单独函数
			defer this.SocketLeave(NewSocketId)
			//停止NewSocketId获取
			break
		}
	}


	// 后续socket的所有消息都在这里执行，如果断开都走defer leave干掉用户，心跳也在这里，目前还不支持多窗口单一心跳，这个将来客户端修改，主要是nginx time out300秒
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		var info server.IM_protocol
		if err := json.Unmarshal([]byte(p), &info); err == nil {
			this.MsgList <- this.NewMsg(info.Type, info.Users, info.SocketId,string(info.Msg))
			//G.Logger.Info(info)

		} else {
			Global.Logger.Error("Join", err)
		}

	}

}

func (this *WebSocketListController)NewMsg(ep server.EventType, user server.IM_protocol_user,SocketId uint32, msg string) server.IM_protocol {
	return server.IM_protocol{ep, msg,SocketId,user, int(time.Now().Unix()) }
}

func (this *WebSocketListController)chatroom() {
	for {
		select {
		case JoinSocket := <-this.SocketChan:
			if !this.IsExistSocketById(JoinSocket.SocketId) {
				this.ActiveSocketList.PushBack(JoinSocket) // Add user to the end of list.
				// Publish a JOIN event.
				this.MsgList <- this.NewMsg(server.IM_EVENT_JOIN, JoinSocket.User,JoinSocket.SocketId,"")
				this.MsgList <- this.NewMsg(server.IM_EVENT_BROADCAST_JOIN, JoinSocket.User,JoinSocket.SocketId,"")
				this.MsgList <- this.NewMsg(server.IM_EVENT_MESSAGE, JoinSocket.User,JoinSocket.SocketId,"welcome")
				Global.Logger.Info("New socket:", JoinSocket.SocketId, ";WebSocket:", JoinSocket.Conn != nil)
			} else {
				Global.Logger.Info("Old socket:", JoinSocket.SocketId, ";WebSocket:", JoinSocket.Conn != nil)
			}
		case SocketMessage := <-this.MsgList:
			//如果是心跳，单发
			switch SocketMessage.Type {
			case
				server.IM_EVENT_HEART,
				server.IM_EVENT_JOIN,
				server.IM_EVENT_LEAVE,
				server.IM_EVENT_MESSAGE:
				this.HeartWebSocket(SocketMessage)
				this.Game(SocketMessage)
				break
			case
				server.IM_EVENT_BROADCAST_HEART,
				server.IM_EVENT_BROADCAST_JOIN,
				server.IM_EVENT_BROADCAST_LEAVE,
				server.IM_EVENT_BROADCAST_MESSAGE:
				this.broadcastWebSocket(SocketMessage)
				break
			}
			server.NewArchive(SocketMessage)
			if SocketMessage.Type == server.IM_EVENT_MESSAGE {
				Global.Logger.Info("Message from", SocketMessage.Users.From, ";Msg:", SocketMessage.Msg)}

		case LeaveSocket := <-this.UnSocketChan:
			for sub := this.ActiveSocketList.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(SocketInfo).SocketId == LeaveSocket.SocketId {
					this.ActiveSocketList.Remove(sub)
					// Clone connection.
					ws := sub.Value.(SocketInfo).Conn
					if ws != nil {
						ws.Close()
						Global.Logger.Error("WebSocket closed:", LeaveSocket)
					}

					this.MsgList  <- this.NewMsg(server.IM_EVENT_LEAVE, sub.Value.(SocketInfo).User,LeaveSocket.SocketId, "") // Publish a LEAVE event.
					this.MsgList  <- this.NewMsg(server.IM_EVENT_BROADCAST_LEAVE, sub.Value.(SocketInfo).User,LeaveSocket.SocketId, "")
					break
				}
			}
		}
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func (this *WebSocketListController)broadcastWebSocket(event server.IM_protocol) {
	data, err := json.Marshal(event)
	if err != nil {
		Global.Logger.Error("Fail to marshal event:", err)
		return
	}

	for sub := this.ActiveSocketList.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(SocketInfo).Conn
		if ws != nil {

			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				Global.Logger.Trace("disconnected user:",sub.Value.(SocketInfo).User)
				this.UnSocketChan <- UnSocketId{sub.Value.(SocketInfo).SocketId}

			}
		}
	}
}

func  (this *WebSocketListController)BCGame(event server.IM_protocol){
	this.broadcastWebSocket(event)
}
func (this *WebSocketListController)NetRussia()  {
	for{
		time.Sleep(400 * time.Millisecond)
		event := server.IM_protocol{}
		event.Type = server.IM_EVENT_BROADCAST_MESSAGE
		ret,b :=Game.Start(event)
		if true==b{
			this.BCGame(ret)
		}
	}
}
func (this *WebSocketListController)Game(event server.IM_protocol)  {
	if ""==event.Msg {
		return
	}
	ret ,t:=Game.Start(event)
	if true==t {
		this.BCGame(ret)
	}

}
func (this *WebSocketListController)HeartWebSocket(event server.IM_protocol) {
	data, err := json.Marshal(event)
	if err != nil {
		Global.Logger.Error("Fail to marshal event:", err)
		return
	}

	for sub := this.ActiveSocketList.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		if sub.Value.(SocketInfo).SocketId == event.SocketId {
			ws := sub.Value.(SocketInfo).Conn
			if ws != nil {
				if ws.WriteMessage(websocket.TextMessage, data) != nil {
					// socket disconnected.
					this.UnSocketChan <- UnSocketId{sub.Value.(SocketInfo).SocketId}
				} else {
					Global.Logger.Trace("Socketheart :",event.SocketId)
				}
			}
		}
	}
}
