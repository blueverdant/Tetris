package Gamecontrollers

import (
	"container/list"
	"github.com/fv0008/AWS_Russia/server"
	"github.com/fv0008/AWS_Russia/server/Global"
	"time"
)

type PrivateChatSocket struct {
	SocketId	 	uint32
	Timestamp 		int // Unix timestamp (secs)
}


type PrivateChatChan struct{
	//private
	PrivateKey      		string
	// Send events here to publish them.
	MsgList 				chan(server.IM_protocol)
	// Socket list.
	ActivePrivateChatSocket *list.List
}


func (this *PrivateChatChan)NewChatMsg(ep server.EventType, user server.IM_protocol_user,Socketid uint32, msg string) server.IM_protocol {
	return server.IM_protocol{ep, msg,Socketid,user, int(time.Now().Unix())}
}


func (this *PrivateChatChan)IsInChatChinById(SocketId uint32) (bool) {
	for sub := this.ActivePrivateChatSocket.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(SocketInfo).SocketId == SocketId {
			return true
		}
	}
	return false
}

func (this *PrivateChatChan) ChatChanAcitve() {
	for {
		select {
		case Message := <-this.MsgList:
			//如果是心跳，单发
			switch Message.Type {
			case
				server.IM_EVENT_HEART,
				server.IM_EVENT_JOIN,
				server.IM_EVENT_LEAVE,
				server.IM_EVENT_MESSAGE:
				//this.HeartWebSocket(event)
				Global.Logger.Info("ChatChinMsg from", Message.Users.From, ";Msg:", Message.Msg)
				break
			case
				server.IM_EVENT_BROADCAST_HEART,
				server.IM_EVENT_BROADCAST_JOIN,
				server.IM_EVENT_BROADCAST_LEAVE,
				server.IM_EVENT_BROADCAST_MESSAGE:
				//this.broadcastWebSocket(event)
				Global.Logger.Info("ChatChinBCMsg from", Message.Users.From, ";Msg:", Message.Msg)
				break
			}
		}
	}
}




func (this *PrivateChatChan) init(){
	// Send events here to publish them.
	this.MsgList = make(chan server.IM_protocol, 10)
	// Long polling waiting list.
	this.ActivePrivateChatSocket = list.New()

	go this.ChatChanAcitve()

}


