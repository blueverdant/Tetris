package Gamecontrollers

import (
	"container/list"
	"github.com/fv0008/AWS_Russia/server"
	"github.com/fv0008/gocket/src/Global"
	"github.com/fv0008/gocket/src/lunarhookmodel"
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


func (this *PrivateChatChan)NewChatMsg(ep lunarhookmodel.EventType, user lunarhookmodel.IM_protocol_user,Socketid uint32, msg string) lunarhookmodel.IM_protocol {
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
				lunarhookmodel.IM_EVENT_HEART,
				lunarhookmodel.IM_EVENT_JOIN,
				lunarhookmodel.IM_EVENT_LEAVE,
				lunarhookmodel.IM_EVENT_MESSAGE:
				//this.HeartWebSocket(event)
				Global.Logger.Info("ChatChinMsg from", Message.Users.From, ";Msg:", Message.Msg)
				break
			case
				lunarhookmodel.IM_EVENT_BROADCAST_HEART,
				lunarhookmodel.IM_EVENT_BROADCAST_JOIN,
				lunarhookmodel.IM_EVENT_BROADCAST_LEAVE,
				lunarhookmodel.IM_EVENT_BROADCAST_MESSAGE:
				//this.broadcastWebSocket(event)
				Global.Logger.Info("ChatChinBCMsg from", Message.Users.From, ";Msg:", Message.Msg)
				break
			}
		}
	}
}




func (this *PrivateChatChan) init(){
	// Send events here to publish them.
	this.MsgList = make(chan lunarhookmodel.IM_protocol, 10)
	// Long polling waiting list.
	this.ActivePrivateChatSocket = list.New()

	go this.ChatChanAcitve()

}


